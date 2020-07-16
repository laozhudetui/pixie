#pragma once

#include <memory>
#include <random>
#include <string>
#include <utility>

#include "src/common/base/base.h"
#include "src/stirling/dynamic_tracing/dynamic_tracer.h"
#include "src/stirling/source_connector.h"

namespace pl {
namespace stirling {

class DynamicTraceConnector : public SourceConnector, public bpf_tools::BCCWrapper {
 public:
  ~DynamicTraceConnector() override = default;

  static StatusOr<std::unique_ptr<SourceConnector>> Create(
      const dynamic_tracing::ir::logical::Program& program) {
    PL_ASSIGN_OR_RETURN(dynamic_tracing::BCCProgram bcc_program,
                        dynamic_tracing::CompileProgram(program));

    if (bcc_program.perf_buffer_specs.size() != 1) {
      return error::Internal("Only a single output table is allowed for now.");
    }

    auto& output = bcc_program.perf_buffer_specs[0];
    PL_ASSIGN_OR_RETURN(std::unique_ptr<DynamicDataTableSchema> table_schema,
                        DynamicDataTableSchema::Create(std::move(output.output)));

    // Make the source connector name the same as the table name.
    // This should be okay so long as there is only one table per connector.
    std::string name(table_schema->Get().name());

    return std::unique_ptr<SourceConnector>(
        new DynamicTraceConnector(name, std::move(table_schema), std::move(bcc_program)));
  }

 protected:
  // TODO(oazizi): This constructor only works with a single table,
  //               since the ArrayView creation only works for a single schema.
  //               Consider how to expand to multiple tables if/when needed.
  DynamicTraceConnector(std::string_view name, std::unique_ptr<DynamicDataTableSchema> table_schema,
                        dynamic_tracing::BCCProgram bcc_program)
      : SourceConnector(name, ArrayView<DataTableSchema>(&table_schema->Get(), 1)),
        table_schema_(std::move(table_schema)),
        bcc_program_(std::move(bcc_program)),
        coin_flip_dist_(0, 1) {}

  Status InitImpl() override;

  void TransferDataImpl(ConnectorContext* ctx, uint32_t table_num, DataTable* data_table) override;

  Status StopImpl() override { return Status::OK(); }

 private:
  std::unique_ptr<DynamicDataTableSchema> table_schema_;
  dynamic_tracing::BCCProgram bcc_program_;

  // TODO(oazizi): Temporary remove.
  std::default_random_engine rng_;
  std::uniform_int_distribution<int> coin_flip_dist_;
};

}  // namespace stirling
}  // namespace pl