#include <gmock/gmock.h>
#include <google/protobuf/text_format.h>
#include <gtest/gtest.h>

#include <utility>
#include <vector>

#include <pypa/parser/parser.hh>

#include "src/carnot/compiler/ir_nodes.h"
#include "src/carnot/compiler/metadata_handler.h"
#include "src/carnot/compiler/physical_planner.h"
#include "src/carnot/compiler/rule_mock.h"
#include "src/carnot/compiler/rules.h"
#include "src/carnot/compiler/test_utils.h"
#include "src/carnot/udf_exporter/udf_exporter.h"
#include "src/common/testing/protobuf.h"

namespace pl {
namespace carnot {
namespace compiler {
namespace physical {
using pl::testing::proto::EqualsProto;

const char* kOneAgentOneKelvinPhysicalState = R"proto(
carnot_info {
  query_broker_address: "agent"
  has_grpc_server: false
  has_data_store: true
  processes_data: true
  accepts_remote_sources: false
}
carnot_info {
  query_broker_address: "kelvin"
  grpc_address: "1111"
  has_grpc_server: true
  has_data_store: false
  processes_data: true
  accepts_remote_sources: true
}
)proto";

class PhysicalPlannerTest : public OperatorTests {
 protected:
  void SetUpImpl() override { compiler_state_ = nullptr; }
  compilerpb::PhysicalState LoadPhysicalStatePb(const std::string& physical_state_txt) {
    compilerpb::PhysicalState physical_state_pb;
    CHECK(google::protobuf::TextFormat::MergeFromString(physical_state_txt, &physical_state_pb));
    return physical_state_pb;
  }

  std::unique_ptr<CompilerState> compiler_state_;
};

const char* kOneAgentOneKelvinPhysicalPlan = R"proto(
qb_address_to_plan {
  key: "agent"
  value {
    dag {
      nodes {
        id: 1
      }
    }
    nodes {
      id: 1
      dag {
        nodes {
          id: 0
          sorted_children: 4
        }
        nodes {
          id: 4
          sorted_parents: 0
        }
      }
      nodes {
        op {
          op_type: MEMORY_SOURCE_OPERATOR
          mem_source_op {
            name: "table"
            column_idxs: 0
            column_idxs: 1
            column_idxs: 2
            column_idxs: 3
            column_names: "count"
            column_names: "cpu0"
            column_names: "cpu1"
            column_names: "cpu2"
            column_types: INT64
            column_types: FLOAT64
            column_types: FLOAT64
            column_types: FLOAT64
          }
        }
      }
      nodes {
        id: 4
        op {
          op_type: GRPC_SINK_OPERATOR
          grpc_sink_op {
            address: "1111"
            destination_id: "agent:0"
          }
        }
      }
    }
  }
}
qb_address_to_plan {
  key: "kelvin"
  value {
    dag {
      nodes {
        id: 1
      }
    }
    nodes {
      id: 1
      dag {
        nodes {
          id: 4
          sorted_children: 2
        }
        nodes {
          id: 2
          sorted_parents: 4
        }
      }
      nodes {
        id: 4
        op {
          op_type: GRPC_SOURCE_OPERATOR
          grpc_source_op {
            source_id: "agent:0"
            column_types: INT64
            column_types: FLOAT64
            column_types: FLOAT64
            column_types: FLOAT64
            column_names: "count"
            column_names: "cpu0"
            column_names: "cpu1"
            column_names: "cpu2"
          }
        }
      }
      nodes {
        id: 2
        op {
          op_type: MEMORY_SINK_OPERATOR
          mem_sink_op {
            name: "out"
            column_types: INT64
            column_types: FLOAT64
            column_types: FLOAT64
            column_types: FLOAT64
            column_names: "count"
            column_names: "cpu0"
            column_names: "cpu1"
            column_names: "cpu2"
          }
        }
      }
    }
  }
}
qb_address_to_dag_id {
  key: "agent"
  value: 1
}
qb_address_to_dag_id {
  key: "kelvin"
  value: 0
}
dag {
  nodes {
    id: 1
    sorted_children: 0
  }
  nodes {
    sorted_parents: 1
  }
}
)proto";

TEST_F(PhysicalPlannerTest, one_agent_one_kelvin) {
  auto mem_src = MakeMemSource(MakeRelation());
  auto mem_sink = MakeMemSink(mem_src, "out");
  PL_CHECK_OK(mem_sink->SetRelation(MakeRelation()));

  compilerpb::PhysicalState ps_pb = LoadPhysicalStatePb(kOneAgentOneKelvinPhysicalState);
  std::unique_ptr<PhysicalPlanner> physical_planner = PhysicalPlanner::Create().ConsumeValueOrDie();
  // TODO(philkuz) fix nullptr for compiler_state.
  std::unique_ptr<PhysicalPlan> physical_plan =
      physical_planner->Plan(ps_pb, compiler_state_.get(), graph.get()).ConsumeValueOrDie();
  EXPECT_THAT(physical_plan->ToProto().ConsumeValueOrDie(),
              EqualsProto(kOneAgentOneKelvinPhysicalPlan));
}

const char* kThreeAgentsOneKelvinPhysicalState = R"proto(
carnot_info {
  query_broker_address: "agent1"
  has_grpc_server: false
  has_data_store: true
  processes_data: true
  accepts_remote_sources: false
}
carnot_info {
  query_broker_address: "agent2"
  has_grpc_server: false
  has_data_store: true
  processes_data: true
  accepts_remote_sources: false
}
carnot_info {
  query_broker_address: "agent3"
  has_grpc_server: false
  has_data_store: true
  processes_data: true
  accepts_remote_sources: false
}
carnot_info {
  query_broker_address: "kelvin"
  grpc_address: "1111"
  has_grpc_server: true
  has_data_store: false
  processes_data: true
  accepts_remote_sources: true
}
)proto";

const char* kThreeAgentsOneKelvinPhysicalPlan = R"proto(
qb_address_to_plan {
  key: "agent1"
  value {
    dag {
      nodes {
        id: 1
      }
    }
    nodes {
      id: 1
      dag {
        nodes {
          sorted_children: 4
        }
        nodes {
          id: 4
          sorted_parents: 0
        }
      }
      nodes {
        op {
          op_type: MEMORY_SOURCE_OPERATOR
          mem_source_op {
            name: "table"
            column_idxs: 0
            column_idxs: 1
            column_idxs: 2
            column_idxs: 3
            column_names: "count"
            column_names: "cpu0"
            column_names: "cpu1"
            column_names: "cpu2"
            column_types: INT64
            column_types: FLOAT64
            column_types: FLOAT64
            column_types: FLOAT64
          }
        }
      }
      nodes {
        id: 4
        op {
          op_type: GRPC_SINK_OPERATOR
          grpc_sink_op {
            address: "1111"
            destination_id: "agent1:0"
          }
        }
      }
    }
  }
}
qb_address_to_plan {
  key: "agent2"
  value {
    dag {
      nodes {
        id: 1
      }
    }
    nodes {
      id: 1
      dag {
        nodes {
          sorted_children: 4
        }
        nodes {
          id: 4
          sorted_parents: 0
        }
      }
      nodes {
        op {
          op_type: MEMORY_SOURCE_OPERATOR
          mem_source_op {
            name: "table"
            column_idxs: 0
            column_idxs: 1
            column_idxs: 2
            column_idxs: 3
            column_names: "count"
            column_names: "cpu0"
            column_names: "cpu1"
            column_names: "cpu2"
            column_types: INT64
            column_types: FLOAT64
            column_types: FLOAT64
            column_types: FLOAT64
          }
        }
      }
      nodes {
        id: 4
        op {
          op_type: GRPC_SINK_OPERATOR
          grpc_sink_op {
            address: "1111"
            destination_id: "agent2:0"
          }
        }
      }
    }
  }
}
qb_address_to_plan {
  key: "agent3"
  value {
    dag {
      nodes {
        id: 1
      }
    }
    nodes {
      id: 1
      dag {
        nodes {
          sorted_children: 4
        }
        nodes {
          id: 4
          sorted_parents: 0
        }
      }
      nodes {
        op {
          op_type: MEMORY_SOURCE_OPERATOR
          mem_source_op {
            name: "table"
            column_idxs: 0
            column_idxs: 1
            column_idxs: 2
            column_idxs: 3
            column_names: "count"
            column_names: "cpu0"
            column_names: "cpu1"
            column_names: "cpu2"
            column_types: INT64
            column_types: FLOAT64
            column_types: FLOAT64
            column_types: FLOAT64
          }
        }
      }
      nodes {
        id: 4
        op {
          op_type: GRPC_SINK_OPERATOR
          grpc_sink_op {
            address: "1111"
            destination_id: "agent3:0"
          }
        }
      }
    }
  }
}
qb_address_to_plan {
  key: "kelvin"
  value {
    dag {
      nodes {
        id: 1
      }
    }
    nodes {
      id: 1
      dag {
        nodes {
          id: 6
          sorted_children: 7
        }
        nodes {
          id: 5
          sorted_children: 7
        }
        nodes {
          id: 4
          sorted_children: 7
        }
        nodes {
          id: 7
          sorted_children: 2
          sorted_parents: 4
          sorted_parents: 5
          sorted_parents: 6
        }
        nodes {
          id: 2
          sorted_parents: 7
        }
      }
      nodes {
        id: 6
        op {
          op_type: GRPC_SOURCE_OPERATOR
          grpc_source_op {
            source_id: "agent1:0"
            column_types: INT64
            column_types: FLOAT64
            column_types: FLOAT64
            column_types: FLOAT64
            column_names: "count"
            column_names: "cpu0"
            column_names: "cpu1"
            column_names: "cpu2"
          }
        }
      }
      nodes {
        id: 5
        op {
          op_type: GRPC_SOURCE_OPERATOR
          grpc_source_op {
            source_id: "agent2:0"
            column_types: INT64
            column_types: FLOAT64
            column_types: FLOAT64
            column_types: FLOAT64
            column_names: "count"
            column_names: "cpu0"
            column_names: "cpu1"
            column_names: "cpu2"
          }
        }
      }
      nodes {
        id: 4
        op {
          op_type: GRPC_SOURCE_OPERATOR
          grpc_source_op {
            source_id: "agent3:0"
            column_types: INT64
            column_types: FLOAT64
            column_types: FLOAT64
            column_types: FLOAT64
            column_names: "count"
            column_names: "cpu0"
            column_names: "cpu1"
            column_names: "cpu2"
          }
        }
      }
      nodes {
        id: 7
        op {
          op_type: UNION_OPERATOR
          union_op {
            column_names: "count"
            column_names: "cpu0"
            column_names: "cpu1"
            column_names: "cpu2"
            column_mappings {
              column_indexes: 0
              column_indexes: 1
              column_indexes: 2
              column_indexes: 3
            }
            column_mappings {
              column_indexes: 0
              column_indexes: 1
              column_indexes: 2
              column_indexes: 3
            }
            column_mappings {
              column_indexes: 0
              column_indexes: 1
              column_indexes: 2
              column_indexes: 3
            }
          }
        }
      }
      nodes {
        id: 2
        op {
          op_type: MEMORY_SINK_OPERATOR
          mem_sink_op {
            name: "out"
            column_types: INT64
            column_types: FLOAT64
            column_types: FLOAT64
            column_types: FLOAT64
            column_names: "count"
            column_names: "cpu0"
            column_names: "cpu1"
            column_names: "cpu2"
          }
        }
      }
    }
  }
}
qb_address_to_dag_id {
  key: "agent1"
  value: 1
}
qb_address_to_dag_id {
  key: "agent2"
  value: 2
}
qb_address_to_dag_id {
  key: "agent3"
  value: 3
}
qb_address_to_dag_id {
  key: "kelvin"
  value: 0
}
dag {
  nodes {
    id: 3
    sorted_children: 0
  }
  nodes {
    id: 2
    sorted_children: 0
  }
  nodes {
    id: 1
    sorted_children: 0
  }
  nodes {
    sorted_parents: 1
    sorted_parents: 2
    sorted_parents: 3
  }
}
)proto";

TEST_F(PhysicalPlannerTest, three_agent_one_kelvin) {
  auto mem_src = MakeMemSource(MakeRelation());
  auto mem_sink = MakeMemSink(mem_src, "out");
  PL_CHECK_OK(mem_sink->SetRelation(MakeRelation()));

  compilerpb::PhysicalState ps_pb = LoadPhysicalStatePb(kThreeAgentsOneKelvinPhysicalState);
  std::unique_ptr<PhysicalPlanner> physical_planner = PhysicalPlanner::Create().ConsumeValueOrDie();
  std::unique_ptr<PhysicalPlan> physical_plan =
      physical_planner->Plan(ps_pb, compiler_state_.get(), graph.get()).ConsumeValueOrDie();
  EXPECT_THAT(physical_plan->ToProto().ConsumeValueOrDie(),
              EqualsProto(kThreeAgentsOneKelvinPhysicalPlan));
}

}  // namespace physical
}  // namespace compiler
}  // namespace carnot
}  // namespace pl
