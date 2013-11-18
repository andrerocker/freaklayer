#!/usr/bin/env ruby
require "pty"

def buildpacks
    @buildpacks ||= Dir["#{File.absolute_path(File.dirname(__FILE__))}/*/"].select do |current|
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
    command "#{detect}/bin/compile #{ARGV.first} #{ARGV.last}"
end

def command(command)
    PTY.spawn(command) do |stdin, stdout, pid|
        stdin.each { |line| print line }
    end rescue nil
end

execute
