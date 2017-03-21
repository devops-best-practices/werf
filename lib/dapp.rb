require 'pathname'
require 'fileutils'
require 'tmpdir'
require 'digest'
require 'timeout'
require 'base64'
require 'mixlib/cli'
require 'mixlib/shellout'
require 'securerandom'
require 'excon'
require 'json'
require 'uri'
require 'ostruct'
require 'time'
require 'i18n'
require 'paint'
require 'inifile'
require 'rugged'
require 'rubygems/package'

require 'net_status'

require 'dapp/version'
require 'dapp/core_ext/hash'
require 'dapp/core_ext/pathname'
require 'dapp/helper/cli'
require 'dapp/helper/trivia'
require 'dapp/helper/sha256'
require 'dapp/helper/net_status'
require 'dapp/prctl'
require 'dapp/error/base'
require 'dapp/error/dapp'
require 'dapp/error/dappfile'
require 'dapp/error/shellout'
require 'dapp/exception/base'
require 'dapp/cli'
require 'dapp/cli/base'
require 'dapp/config/base'
require 'dapp/config/directive/base'
require 'dapp/config/error/config'
require 'dapp/dapp/lock'
require 'dapp/dapp/ssh_agent'
require 'dapp/dapp/git_artifact'
require 'dapp/dapp/dappfile'
require 'dapp/dapp/chef'
require 'dapp/dapp/logging/base'
require 'dapp/dapp/logging/process'
require 'dapp/dapp/logging/paint'
require 'dapp/dapp/logging/i18n'
require 'dapp/dapp/deps/base'
require 'dapp/dapp/deps/gitartifact'
require 'dapp/dapp/shellout/streaming'
require 'dapp/dapp/shellout/base'
require 'dapp/dapp'
require 'dapp/dimg'
require 'dapp/dimg/builder'
require 'dapp/dimg/builder/base'
require 'dapp/dimg/builder/chef'
require 'dapp/dimg/builder/chef/cookbook_metadata'
require 'dapp/dimg/builder/chef/berksfile'
require 'dapp/dimg/builder/chef/cookbook'
require 'dapp/dimg/builder/shell'
require 'dapp/dimg/builder/none'
require 'dapp/dimg/build/stage/mod/logging'
require 'dapp/dimg/build/stage/mod/group'
require 'dapp/dimg/build/stage/mod/git_artifact_dependencies'
require 'dapp/dimg/build/stage/base'
require 'dapp/dimg/build/stage/ga_base'
require 'dapp/dimg/build/stage/ga_dependencies_base'
require 'dapp/dimg/build/stage/artifact_base'
require 'dapp/dimg/build/stage/artifact_default'
require 'dapp/dimg/build/stage/from'
require 'dapp/dimg/build/stage/import_artifact'
require 'dapp/dimg/build/stage/before_install'
require 'dapp/dimg/build/stage/before_install_artifact'
require 'dapp/dimg/build/stage/ga_archive_dependencies'
require 'dapp/dimg/build/stage/ga_archive'
require 'dapp/dimg/build/stage/install/ga_pre_install_patch_dependencies'
require 'dapp/dimg/build/stage/install/ga_pre_install_patch'
require 'dapp/dimg/build/stage/install/install'
require 'dapp/dimg/build/stage/install/ga_post_install_patch_dependencies'
require 'dapp/dimg/build/stage/install/ga_post_install_patch'
require 'dapp/dimg/build/stage/after_install_artifact'
require 'dapp/dimg/build/stage/before_setup'
require 'dapp/dimg/build/stage/before_setup_artifact'
require 'dapp/dimg/build/stage/setup/ga_pre_setup_patch_dependencies'
require 'dapp/dimg/build/stage/setup/ga_pre_setup_patch'
require 'dapp/dimg/build/stage/setup/setup'
require 'dapp/dimg/build/stage/setup/ga_post_setup_patch_dependencies'
require 'dapp/dimg/build/stage/setup/ga_post_setup_patch'
require 'dapp/dimg/build/stage/after_setup_artifact'
require 'dapp/dimg/build/stage/ga_latest_patch'
require 'dapp/dimg/build/stage/docker_instructions'
require 'dapp/dimg/build/stage/ga_artifact_patch'
require 'dapp/dimg/build/stage/build_artifact'
require 'dapp/dimg/cli/base'
require 'dapp/dimg/cli/dimg'
require 'dapp/dimg/cli/build'
require 'dapp/dimg/cli/push'
require 'dapp/dimg/cli/spush'
require 'dapp/dimg/cli/tag'
require 'dapp/dimg/cli/list'
require 'dapp/dimg/cli/build_context'
require 'dapp/dimg/cli/build_context/export'
require 'dapp/dimg/cli/build_context/import'
require 'dapp/dimg/cli/stages'
require 'dapp/dimg/cli/stages/flush_local'
require 'dapp/dimg/cli/stages/flush_repo'
require 'dapp/dimg/cli/stages/cleanup_local'
require 'dapp/dimg/cli/stages/cleanup_repo'
require 'dapp/dimg/cli/stages/push'
require 'dapp/dimg/cli/stages/pull'
require 'dapp/dimg/cli/run'
require 'dapp/dimg/cli/cleanup'
require 'dapp/dimg/cli/bp'
require 'dapp/dimg/cli/mrproper'
require 'dapp/dimg/cli/stage_image'
require 'dapp/dimg/cli/cli'
require 'dapp/dimg/config/base'
require 'dapp/dimg/config/directive/base'
require 'dapp/dimg/config/dimg/instance_methods'
require 'dapp/dimg/config/dimg/validation'
require 'dapp/dimg/config/dimg'
require 'dapp/dimg/config/artifact_dimg'
require 'dapp/dimg/config/dimg_group_base'
require 'dapp/dimg/config/dimg_group_main'
require 'dapp/dimg/config/dimg_group'
require 'dapp/dimg/config/artifact_group'
require 'dapp/dimg/config/directive/artifact_base'
require 'dapp/dimg/config/directive/git_artifact_local'
require 'dapp/dimg/config/directive/git_artifact_remote'
require 'dapp/dimg/config/directive/artifact'
require 'dapp/dimg/config/directive/chef'
require 'dapp/dimg/config/directive/shell/dimg'
require 'dapp/dimg/config/directive/shell/artifact'
require 'dapp/dimg/config/directive/docker/base'
require 'dapp/dimg/config/directive/docker/dimg'
require 'dapp/dimg/config/directive/docker/artifact'
require 'dapp/dimg/config/directive/mount'
require 'dapp/dimg/dapp/command/common'
require 'dapp/dimg/dapp/command/build'
require 'dapp/dimg/dapp/command/bp'
require 'dapp/dimg/dapp/command/cleanup'
require 'dapp/dimg/dapp/command/mrproper'
require 'dapp/dimg/dapp/command/stage_image'
require 'dapp/dimg/dapp/command/list'
require 'dapp/dimg/dapp/command/push'
require 'dapp/dimg/dapp/command/run'
require 'dapp/dimg/dapp/command/spush'
require 'dapp/dimg/dapp/command/tag'
require 'dapp/dimg/dapp/command/stages/common'
require 'dapp/dimg/dapp/command/stages/cleanup_local'
require 'dapp/dimg/dapp/command/stages/cleanup_repo'
require 'dapp/dimg/dapp/command/stages/flush_local'
require 'dapp/dimg/dapp/command/stages/flush_repo'
require 'dapp/dimg/dapp/command/stages/push'
require 'dapp/dimg/dapp/command/stages/pull'
require 'dapp/dimg/dapp/command/build_context/export'
require 'dapp/dimg/dapp/command/build_context/import'
require 'dapp/dimg/dapp/command/build_context/common'
require 'dapp/dimg/dapp/command'
require 'dapp/dimg/docker_registry'
require 'dapp/dimg/docker_registry/base/request'
require 'dapp/dimg/docker_registry/base/authorization'
require 'dapp/dimg/docker_registry/base'
require 'dapp/dimg/docker_registry/default'
require 'dapp/dimg/error/base'
require 'dapp/dimg/error/dimg'
require 'dapp/dimg/error/build'
require 'dapp/dimg/error/tar_writer'
require 'dapp/dimg/error/rugged'
require 'dapp/dimg/error/registry'
require 'dapp/dimg/error/chef'
require 'dapp/dimg/error/lock'
require 'dapp/dimg/error/config'
require 'dapp/dimg/error/command'
require 'dapp/dimg/exception/base'
require 'dapp/dimg/exception/introspect_image'
require 'dapp/dimg/exception/registry'
require 'dapp/dimg/lock/base'
require 'dapp/dimg/lock/file'
require 'dapp/dimg/filelock'
require 'dapp/dimg/git_repo/base'
require 'dapp/dimg/git_repo/own'
require 'dapp/dimg/git_repo/remote'
require 'dapp/dimg/git_artifact'
require 'dapp/dimg/image/argument'
require 'dapp/dimg/image/docker'
require 'dapp/dimg/image/stage'
require 'dapp/dimg/image/scratch'
require 'dapp/dimg/dimg/git_artifact'
require 'dapp/dimg/dimg/path'
require 'dapp/dimg/dimg/tags'
require 'dapp/dimg/dimg/stages'
require 'dapp/dimg/dimg'
require 'dapp/dimg/artifact'

module Dapp
  def self.root
    File.expand_path('../..', __FILE__)
  end
end
