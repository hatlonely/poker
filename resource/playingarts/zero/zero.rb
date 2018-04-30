#!/usr/bin/env ruby

require 'uri'
require "open-uri"

File.open("zero.txt", "r") do |file|
    file.each_line do |line|
        name = URI.parse(line).path.split('/').last
        puts line
        open(line) {|data|
            File.open(name, "wb") do |jpg|
              jpg.puts data.read
            end
         }
    end
end