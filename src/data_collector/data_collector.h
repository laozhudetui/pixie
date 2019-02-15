#pragma once

#include <chrono>
#include <functional>
#include <memory>
#include <string>
#include <thread>
#include <unordered_map>
#include <utility>
#include <vector>

#include "src/common/error.h"
#include "src/common/status.h"
#include "src/data_collector/data_table.h"
#include "src/data_collector/pub_sub_manager.h"
#include "src/data_collector/source_connector.h"
#include "src/data_collector/source_registry.h"

namespace pl {
namespace datacollector {

/**
 * The data collector collects data from various different 'sources',
 * and makes them available via a structured API, where the data can then be used and queried as
 * needed (by Pixie or others). Its function is to unify various, disparate sources of data into a
 * common, structured data format.
 */
class DataCollector {
 public:
  DataCollector() = delete;
  explicit DataCollector(std::unique_ptr<SourceRegistry> registry)
      : registry_(std::move(registry)) {
    config_ = std::make_unique<PubSubManager>(schemas_);
  }
  ~DataCollector() = default;

  /**
   * @brief Create data source connectors from the registered sources.
   *
   * @return Status
   */
  Status CreateSourceConnectors();

  /**
   * Register call-back from Agent. Used to periodically send data.
   *
   * Function signature is:
   *   uint64_t table_id
   *   std::unique_ptr<ColumnWrapperRecordBatch> data
   */
  void RegisterCallback(
      std::function<void(uint64_t, std::unique_ptr<ColumnWrapperRecordBatch>)> f) {
    agent_callback_ = f;
  }

  /**
   * Main data collection call, that is spawned off as an independent thread.
   */
  void Run();

  /**
   * Wait the running thread to terminate.
   */
  void Wait();

 private:
  /**
   * Main data source polling loop.
   */
  void RunThread();

  /**
   * Main data source polling loop.
   */
  Status AddSource(const std::string& name, std::unique_ptr<SourceConnector> source);

  /**
   * Helper function to figure out how much to sleep between polling iterations.
   */
  void SleepUntilNextTick();

  /**
   * Main thread used to spawn off RunThread().
   */
  std::thread run_thread_;

  /**
   * Vector of all Source Connectors.
   */
  std::vector<std::unique_ptr<SourceConnector>> sources_;

  /**
   * Vector of all Data Tables.
   */
  std::vector<std::unique_ptr<DataTable>> tables_;

  /**
   * Vector of all the Schemas.
   */
  std::vector<std::unique_ptr<InfoClassSchema>> schemas_;

  /**
   * Pointer the config unit that handles sub/pub with agent.
   */
  std::unique_ptr<PubSubManager> config_;

  /**
   * @brief Pointer to data source registry
   *
   */
  std::unique_ptr<SourceRegistry> registry_;

  // Defining a constant for data collector wrapper testing.
  const std::chrono::milliseconds kDefaultSamplingPeriod{100};

  /**
   * Function to call to push data to the agent.
   * Function signature is:
   *   uint64_t table_id
   *   std::unique_ptr<ColumnWrapperRecordBatch> data
   */
  std::function<void(uint64_t, std::unique_ptr<ColumnWrapperRecordBatch>)> agent_callback_;
};

}  // namespace datacollector
}  // namespace pl
