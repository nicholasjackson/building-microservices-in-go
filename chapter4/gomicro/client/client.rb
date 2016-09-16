#!/usr/bin/env ruby

this_dir = File.expand_path(File.dirname(__FILE__))
lib_dir = File.join(File.dirname(this_dir), 'proto')
$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)

require 'grpc'
require 'multi_json'
require 'kittens_pb'
require 'net/http'
require 'uri'

include Bmigo::Micro

def main()
    puts "Connecting to Kittenserver"
    http = Net::HTTP.new('kittenserver_kittenserver_1',8091)
    request = Net::HTTP::Post.new('/')
    request.content_type = 'application/octet-stream'

    # Create envelope
    envelope = RequestEnvelope.new(service_method: "Kittens.List", seq: 2)
    env = RequestEnvelope.encode(envelope)
    envSize = Array(env.length).pack('N*') 

    # Create message
    message = Request.new(name: "World")
    mess = Request.encode(message)
    messSize = Array(mess.length).pack('N*')

    request.body = envSize + env + messSize + mess
    response = http.request(request)

    puts response.body.inspect
    body = response.body
    
    sizeByteLength = 4
    endPos = 3
   
    len = body[0..endPos].unpack('N*')[0]
    
    endPos = endPos + len
    resp = body[sizeByteLength..endPos]
    envelope = ResponseEnvelope.decode(resp)
   
    puts envelope.inspect
    
    len = body[endPos+1..endPos+sizeByteLength].unpack('N*')[0]
    
    endPos = endPos + sizeByteLength
    resp = body[endPos+1..endPos+len] 
    message = Response.decode(resp)
   
    puts message.inspect
end

main
