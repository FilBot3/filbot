#!/usr/bin/env ruby
#
#
#

desc 'Download filbot dependencies'
task :deps do 
  sh 'go get -u -v github.com/go-chat-bot/bot/...'  
end

desc 'Build the echo plugin'
task :echo => ['deps'] do 
  puts '=> Building echo plugin'
  Dir.chdir('./echo/')
  sh 'go install'
  Dir.chdir('../')
end

desc 'Build the hello_world plugin'
task :helloWorld => ['deps'] do 
  puts '=> Building hello_world plugin'
  Dir.chdir('./hello_world/')
  sh 'go install'
  Dir.chdir('../')
end

desc 'Build filbot and install it.'
task :filbot => ['echo', 'helloWorld'] do 
  puts '=> Building filbot'
  sh 'go install'
end

desc 'Perform all tasks and build.'
task :build => ['deps', 'echo', 'helloWorld', 'filbot']

task :default => ['build']

desc 'Run Filbot'
task :run => ['build'] do 
  puts '=> Running filbot'
  sh 'run_filbot.sh'
end

