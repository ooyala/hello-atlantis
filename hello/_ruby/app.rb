#!/usr/bin/env ruby

require 'sinatra'

enable :logging
set :port, ENV["HTTP_PORT"] || 9876
set :bind, "0.0.0.0"

get "/healthz" do
  "OK\n"
end

get // do
  "Goodbye #{request.path_info} from ruby #{RUBY_VERSION}!"
end
