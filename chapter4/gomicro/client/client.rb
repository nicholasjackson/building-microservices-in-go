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

# read the response from the go micro server
# body: the response returned from the http client
# message_class: the class type expected to be returned in the response
def read_response(body, message_class)
  envelope, envelope_size = read_envelope(body) 

  start_pos = MESSAGE_SIZE_BYTES + envelope_size
  message = read_message(start_pos, body, message_class)

  return envelope, message
end

# reads the message envelope and returns a ResponseEnvelope
def read_envelope(body)

  end_pos = MESSAGE_SIZE_BYTES - 1 
  message_len = body[0..end_pos].unpack('N*')[0]
  
  end_pos = end_pos + message_len
  resp = body[MESSAGE_SIZE_BYTES..end_pos]

  return ResponseEnvelope.decode(resp), message_len
end

# reads the message object from the response stream
def read_message(start_pos, body, message_class)

  end_pos = start_pos + MESSAGE_SIZE_BYTES
  message_len = body[start_pos..end_pos].unpack('N*')[0]
 
  start_pos = end_pos
  end_pos = start_pos + message_len
  message = body[start_pos..end_pos]
  
  message_class.decode(message)
end

# construct the request objects and send them to the server
# server: the address of the server to connect to
# port: the port of the server to connect to
# method: the method to call on the remote endpoint
# request: an instance of a request class to send to with the request
def send_request(server, port, method, request)

  http = Net::HTTP.new(server,port)
  http_request = Net::HTTP::Post.new('/')
  http_request.content_type = 'application/octet-stream'

  # Create envelope
  envelope = RequestEnvelope.new(service_method: method, seq: 2)
  encoded_envelope = RequestEnvelope.encode(envelope)
  # Message size is encoded into 4 bytes as an unsigned integer
  envelope_size = Array(encoded_envelope.length).pack('N*') 

  # Create message
  encoded_message = request.class.encode(request)
  message_size = Array(encoded_message.length).pack('N*')

  # Add the message body in the format of message length followed by the message
  http_request.body = envelope_size + encoded_envelope + message_size + encoded_message
  response = http.request(http_request)

  return response
end

def main()

  puts "Connecting to Kittenserver"

  request = Bmigo::Micro::Request.new(name: 'Nic')
  body = send_request('kittenserver_kittenserver_1', '8091', 'Kittens.Hello', request).body
  envelope, message = read_response(body, Bmigo::Micro::Response)    
  
  puts envelope.inspect 
  puts message.inspect
end

main
