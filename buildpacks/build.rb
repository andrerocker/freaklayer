#!/usr/bin/env ruby

def buildpacks
    @buildpacks ||= Dir["#{File.absolute_path(".")}/*/"].select do |current|
        File.exist?("#{current}/bin/detect") &&
            File.exist?("#{current}/bin/compile")
    end
end

def detect
    @detect ||= buildpacks.detect do |current|
        %x(#{current}/bin/detect #{ARGV.first}).chomp != "no"
    end
end

def execute
    if detect == nil
        puts "no buildpack detected for your project! falohh!"
        return
    end

    puts "Buildpack detected for your project!! -> #{detect}"
    %x(#{detect}/bin/compile #{ARGV.first})
end

execute
