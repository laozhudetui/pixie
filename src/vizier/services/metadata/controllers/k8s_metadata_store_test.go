package controllers

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"testing"
	"time"

	"github.com/cockroachdb/pebble"
	"github.com/cockroachdb/pebble/vfs"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/assert"

	metadatapb "pixielabs.ai/pixielabs/src/shared/k8s/metadatapb"
	storepb "pixielabs.ai/pixielabs/src/vizier/services/metadata/storepb"
	"pixielabs.ai/pixielabs/src/vizier/utils/datastore/pebbledb"
)

func setupMDSTest(t *testing.T) (*pebbledb.DataStore, *MetadataDatastore, func() error) {
	memFS := vfs.NewMem()
	c, err := pebble.Open("test", &pebble.Options{
		FS: memFS,
	})
	if err != nil {
		t.Fatal("failed to initialize a pebbledb")
		os.Exit(1)
	}

	db := pebbledb.New(c, 3*time.Second)
	ts := NewMetadataDatastore(db)
	cleanup := db.Close

	return db, ts, cleanup
}

func TestMetadataDatastore_AddFullResourceUpdate(t *testing.T) {
	db, mds, cleanup := setupMDSTest(t)
	defer cleanup()

	update := &storepb.K8SResource{
		Resource: &storepb.K8SResource_Namespace{
			Namespace: &metadatapb.Namespace{
				Metadata: &metadatapb.ObjectMetadata{
					Name:            "object_md",
					UID:             "ijkl",
					ResourceVersion: "1",
					ClusterName:     "a_cluster",
					OwnerReferences: []*metadatapb.OwnerReference{
						&metadatapb.OwnerReference{
							Kind: "pod",
							Name: "test",
							UID:  "abcd",
						},
					},
					CreationTimestampNS: 4,
					DeletionTimestampNS: 6,
				},
			},
		},
	}

	err := mds.AddFullResourceUpdate(int64(15), update)
	assert.Nil(t, err)

	savedResourceUpdate, err := db.Get(path.Join(fullResourceUpdatePrefix, "00000000000000000015"))
	assert.Nil(t, err)
	savedResourceUpdatePb := &storepb.K8SResource{}
	err = proto.Unmarshal(savedResourceUpdate, savedResourceUpdatePb)
	assert.Nil(t, err)
	assert.Equal(t, update, savedResourceUpdatePb)
}

func TestMetadataDatastore_AddResourceUpdateForTopic(t *testing.T) {
	db, mds, cleanup := setupMDSTest(t)
	defer cleanup()

	update := &storepb.K8SResourceUpdate{
		Update: &metadatapb.ResourceUpdate{
			UpdateVersion: 2,
			Update: &metadatapb.ResourceUpdate_NamespaceUpdate{
				NamespaceUpdate: &metadatapb.NamespaceUpdate{
					UID:              "ijkl",
					Name:             "object_md",
					StartTimestampNS: 4,
					StopTimestampNS:  6,
				},
			},
		},
	}

	err := mds.AddResourceUpdateForTopic(int64(15), "127.0.0.1", update)
	assert.Nil(t, err)

	savedResourceUpdate, err := db.Get(path.Join(topicResourceUpdatePrefix, "127.0.0.1", "00000000000000000015"))
	assert.Nil(t, err)
	savedResourceUpdatePb := &storepb.K8SResourceUpdate{}
	err = proto.Unmarshal(savedResourceUpdate, savedResourceUpdatePb)
	assert.Nil(t, err)
	assert.Equal(t, update, savedResourceUpdatePb)
}

func TestMetadataDatastore_AddResourceUpdate(t *testing.T) {
	db, mds, cleanup := setupMDSTest(t)
	defer cleanup()

	update := &storepb.K8SResourceUpdate{
		Update: &metadatapb.ResourceUpdate{
			UpdateVersion: 2,
			Update: &metadatapb.ResourceUpdate_NamespaceUpdate{
				NamespaceUpdate: &metadatapb.NamespaceUpdate{
					UID:              "ijkl",
					Name:             "object_md",
					StartTimestampNS: 4,
					StopTimestampNS:  6,
				},
			},
		},
	}

	err := mds.AddResourceUpdate(int64(15), update)
	assert.Nil(t, err)

	savedResourceUpdate, err := db.Get(path.Join(topicResourceUpdatePrefix, unscopedTopic, "00000000000000000015"))
	assert.Nil(t, err)
	savedResourceUpdatePb := &storepb.K8SResourceUpdate{}
	err = proto.Unmarshal(savedResourceUpdate, savedResourceUpdatePb)
	assert.Nil(t, err)
	assert.Equal(t, update, savedResourceUpdatePb)
}

func TestMetadataDatastore_FetchResourceUpdates(t *testing.T) {
	tests := []struct {
		name                  string
		from                  int
		to                    int
		topicSpecificVersions []int
		unscopedVersions      []int
		fetchedVersions       []int
	}{
		{
			name:                  "mix1",
			topicSpecificVersions: []int{5, 6, 8, 20, 53, 56},
			unscopedVersions:      []int{2, 4, 7, 12, 13, 40},
			fetchedVersions:       []int{2, 4, 5, 6, 7, 8, 12, 13, 20, 40},
			from:                  0,
			to:                    53,
		},
		{
			name:                  "mix2",
			topicSpecificVersions: []int{5, 6, 8, 20, 53, 56},
			unscopedVersions:      []int{2, 4, 7, 12, 13, 40},
			fetchedVersions:       []int{20, 40, 53, 56},
			from:                  14,
			to:                    57,
		},
		{
			name:                  "equal",
			topicSpecificVersions: []int{4, 5, 7, 8, 10},
			unscopedVersions:      []int{6, 8, 11},
			fetchedVersions:       []int{5, 6, 7, 8, 10, 11},
			from:                  5,
			to:                    12,
		},
		{
			name:                  "topic empty",
			topicSpecificVersions: []int{},
			unscopedVersions:      []int{2, 4, 6, 8, 10},
			fetchedVersions:       []int{4, 6, 8},
			from:                  4,
			to:                    10,
		},
		{
			name:                  "unscoped empty",
			topicSpecificVersions: []int{2, 4, 6, 8, 10},
			unscopedVersions:      []int{},
			fetchedVersions:       []int{4, 6, 8},
			from:                  4,
			to:                    10,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			db, mds, cleanup := setupMDSTest(t)
			defer cleanup()

			// Insert updates into db.
			for _, u := range tc.topicSpecificVersions {
				update := &storepb.K8SResourceUpdate{
					Update: &metadatapb.ResourceUpdate{
						UpdateVersion: int64(u),
					},
				}
				val, err := update.Marshal()
				assert.Nil(t, err)

				db.Set(path.Join(topicResourceUpdatePrefix, "127.0.0.1", fmt.Sprintf("%020d", u)), string(val))
			}
			for _, u := range tc.unscopedVersions {
				update := &storepb.K8SResourceUpdate{
					Update: &metadatapb.ResourceUpdate{
						UpdateVersion: int64(u),
					},
				}
				val, err := update.Marshal()
				assert.Nil(t, err)

				db.Set(path.Join(topicResourceUpdatePrefix, unscopedTopic, fmt.Sprintf("%020d", u)), string(val))
			}

			updates, err := mds.FetchResourceUpdates("127.0.0.1", int64(tc.from), int64(tc.to))
			assert.Nil(t, err)
			assert.Equal(t, len(tc.fetchedVersions), len(updates))

			for i, v := range tc.fetchedVersions {
				assert.Equal(t, int64(v), updates[i].Update.UpdateVersion)
			}
		})
	}
}

func TestMetadataDatastore_GetUpdateVersion(t *testing.T) {
	db, mds, cleanup := setupMDSTest(t)
	defer cleanup()

	db.Set(path.Join(topicVersionPrefix, "127.0.0.1"), "57")
	version, err := mds.GetUpdateVersion("127.0.0.1")
	assert.Nil(t, err)
	assert.Equal(t, int64(57), version)
}

func TestMetadataDatastore_SetUpdateVersion(t *testing.T) {
	db, mds, cleanup := setupMDSTest(t)
	defer cleanup()

	err := mds.SetUpdateVersion("127.0.0.1", 123)
	assert.Nil(t, err)

	savedVersion, err := db.Get(path.Join(topicVersionPrefix, "127.0.0.1"))
	assert.Nil(t, err)
	i, err := strconv.ParseInt(string(savedVersion), 10, 64)
	assert.Nil(t, err)
	assert.Equal(t, int64(123), i)
}