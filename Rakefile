# coding: utf-8

PROJECT_HOME    = File.expand_path(File.dirname(__FILE__))
PROJECT_NAME    = "command-wrap-notify"
BUILD_DIR       = "#{PROJECT_HOME}/build"
GOOS            = "$(go env GOOS)"
GOARCH          = "$(go env GOARCH)"

task :default => %w( build)

desc "production build #{PROJECT_NAME}"
task :build do
  build_target_dir = "#{PROJECT_HOME}/cmd"
  Dir.each_child(build_target_dir) do |s|
    sh <<-SHELL
      go build \
        -o "#{BUILD_DIR}/#{GOOS}_#{GOARCH}/#{s}" \
        #{build_target_dir}/#{s}
    SHELL
    puts "comple ok! :: #{s}"
  end
end

