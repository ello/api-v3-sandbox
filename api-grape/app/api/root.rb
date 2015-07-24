require 'grape'
require_relative './v1/root'

module API
  class Root < Grape::API
    prefix 'api'
    mount API::V1::Root
  end 
end
