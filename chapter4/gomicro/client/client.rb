#!/usr/bin/env ruby

this_dir = File.expand_path(File.dirname(__FILE__))
lib_dir = File.join(File.dirname(this_dir), 'proto')
$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)

require 'grpc'
require 'multi_json'
require 'kittens_services_pb'

include Bmigo::Micro

def main()
  stub = Kittens::Stub.new('kittenserver_kittenserver_1:8091', :this_channel_is_insecure)
  
  message = stub.list(Request.new(name: "World")).message
end

main
