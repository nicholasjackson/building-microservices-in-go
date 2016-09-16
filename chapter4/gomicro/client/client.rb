#!/usr/bin/env ruby

this_dir = File.expand_path(File.dirname(__FILE__))
lib_dir = File.join(File.dirname(this_dir), 'proto')
$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)

require 'grpc'
require 'multi_json'
require 'kittens_services_pb'
require 'net/http'
require 'uri'
require 'google/protobuf'

include Bmigo::Micro

class Kittenserver
    include Google::Protobuf

    optional :string, :service_method, 1
    optional :string, :seq, 2
end

def main()
    # Create request
    kittenserver = Kittenserver.new
    kittenserver.service_method = "something"
    kittenserver.seq = "1234fddfd" 

    puts "Connecting to Kittenserver"
    http = Net::HTTP.new('kittenserver_kittenserver_1',8091)
    request = Net::HTTP::Post.new('/rpc')
    request.content_type = 'application/json'

    message = Request.new(name: "World")
    request.body = Kittenserver.encode_json(kittenserver)
    request.body = request.body + request.encode_json(message)
    puts request.body
    response = http.request(request)
    puts response.body
end

main
