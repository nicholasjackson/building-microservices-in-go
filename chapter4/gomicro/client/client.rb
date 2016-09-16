#!/usr/bin/env ruby

this_dir = File.expand_path(File.dirname(__FILE__))
lib_dir = File.join(File.dirname(this_dir), 'proto')
$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)

require 'grpc'
require 'multi_json'
require 'kittens_services_pb'
require 'net/http'
require 'uri'

include Bmigo::Micro

def main()
    puts "Connecting to Kittenserver"
    http = Net::HTTP.new('kittenserver_kittenserver_1',8091)
    request = Net::HTTP::Post.new('/rpc')
    request.content_type = 'application/json'

    message = Request.new(name: "World")
    request.body = Request.encode_json(message)
    puts request.body
    response = http.request(request)
    puts response.body
end

main
