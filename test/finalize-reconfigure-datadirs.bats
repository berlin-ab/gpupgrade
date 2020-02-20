#!/usr/bin/env bats

load helpers

setup() {
    skip_if_no_gpdb

    [ ! -z $GPHOME ]
    GPHOME_NEW=${GPHOME_NEW:-$GPHOME}
    GPHOME_OLD=$GPHOME

    setup_state_dir

    gpupgrade kill-services
}

teardown() {
    print_teardown_banner
    teardown_new_cluster
    gpupgrade kill-services

    # reload old path and start
    source "${GPHOME_OLD}/greenplum_path.sh"
    gpstart -a
}

@test "it swaps out the target cluster's data directories and archives the source cluster's data directories" {
    log "Using state directory: $STATE_DIR"

    # place marker file in source master data directory
    local marker_file=source-cluster.test-marker
    touch "$MASTER_DATA_DIRECTORY/${marker_file}"

    log "initialize"
    gpupgrade initialize \
        --old-bindir="$GPHOME/bin" \
        --new-bindir="$GPHOME_NEW/bin" \
        --old-port="${PGPORT}" \
        --link \
        --disk-free-ratio 0 \
        --verbose

    log "execute"
    gpupgrade execute --verbose

    log "finalize"
    gpupgrade finalize

    # ensure source cluster has been archived and target cluster is located where the source used to be
    local source_cluster_master_data_directory="${MASTER_DATA_DIRECTORY}_old"
    local target_cluster_master_data_directory="${MASTER_DATA_DIRECTORY}"
    [ -f "${source_cluster_master_data_directory}/${marker_file}" ] || fail "expected ${marker_file} marker file to be in source datadir: ${STATE_DIR}/base/demoDataDir-1"
    [ ! -f "${target_cluster_master_data_directory}/${marker_file}" ] || fail "unexpected ${marker_file} marker file in target datadir: ${STATE_DIR}/base/demoDataDir-1"

    # ensure gpperfmon configuration file has been modified to reflect new data dir location
    local gpperfmon_config_file="${target_cluster_master_data_directory}/gpperfmon/conf/gpperfmon.conf"
    grep "${target_cluster_master_data_directory}" "${gpperfmon_config_file}" || \
        fail "got gpperfmon.conf file $(cat $gpperfmon_config_file), wanted it to include ${target_cluster_master_data_directory}"

    # ensure that the new cluster is queryable, and has updated configuration
    segment_configuration=$($GPHOME_NEW/bin/psql --no-align --tuples-only postgres -c "select *, version() from gp_segment_configuration")
    [[ $segment_configuration == *"$target_cluster_master_data_directory"* ]] || fail "expected $segment_configuration to include $target_cluster_master_data_directory"
}

setup_state_dir() {
    STATE_DIR=$(mktemp -d /tmp/gpupgrade.XXXXXX)
    export GPUPGRADE_HOME="${STATE_DIR}/gpupgrade"
}

teardown_new_cluster() {
    delete_finalized_cluster $MASTER_DATA_DIRECTORY
}
