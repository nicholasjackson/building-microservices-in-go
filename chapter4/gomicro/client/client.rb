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

MESSAGE_SIZE_BYTES = 4

def read_response(body)

  envelope, envelope_size = read_envelope(body) 

  start_pos = MESSAGE_SIZE_BYTES + envelope_size
  message = read_message(start_pos, body)

  return envelope, message
end

def read_envelope(body)
  end_pos = MESSAGE_SIZE_BYTES - 1
 
  message_len = body[0..end_pos].unpack('N*')[0]
  
  end_pos = end_pos + message_len
  resp = body[MESSAGE_SIZE_BYTES..end_pos]

  return ResponseEnvelope.decode(resp), message_len
end

def read_message(start_pos, body)
  end_pos = start_pos + MESSAGE_SIZE_BYTES
  message_len = body[start_pos..end_pos].unpack('N*')[0]
  start_pos = end_pos
  end_pos = start_pos + message_len
  message = body[start_pos..end_pos]
  puts message.inspect  
  Response.decode(message)
end

def send_request()

  http = Net::HTTP.new('kittenserver_kittenserver_1',8091)
  request = Net::HTTP::Post.new('/')
  request.content_type = 'application/octet-stream'

  # Create envelope
  envelope = RequestEnvelope.new(service_method: "Kittens.List", seq: 2)
  encoded_envelope = RequestEnvelope.encode(envelope)
  envelope_size = Array(encoded_envelope.length).pack('N*') 

  # Create message
  message = Request.new(name: "World")
  encoded_message = Request.encode(message)
  message_size = Array(encoded_message.length).pack('N*')

  request.body = envelope_size + encoded_envelope + message_size + encoded_message
  response = http.request(request)

  return response
end

def main()

    puts "Connecting to Kittenserver"

    body = send_request().body
    envelope, message = read_response(body)    
    
    puts envelope.inspect 
    puts message.inspect
end

main
