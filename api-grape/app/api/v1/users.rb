require 'grape'
require 'json'

module API
	module V1
		class Users < Grape::API
			version 'v1'
			format :json

			resource :users do 
				desc "Returns a list of users"
				get do
					[
						{:username => 'rtyer', :name => 'Ryan Tyer'},
						{:username => 'jayzes', :name => 'Jay Zeschin'}
					].to_json
				end
			end
		end
	end
end