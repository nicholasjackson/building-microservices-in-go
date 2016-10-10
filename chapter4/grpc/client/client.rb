#!/usr/bin/env ruby
this_dir = File.expand_path(File.dirname(__FILE__))
lib_dir = File.join(File.dirname(this_dir), 'proto')
$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)

require 'kittens_services_pb'

service = Bmigo::Grpc::Kittens::Stub.new('kittenserver_kittenserver_1:9000', :this_channel_is_insecure)

request = Bmigo::Grpc::Request.new
request.name = 'Nic'

response = service.hello(request)

puts response.msg
