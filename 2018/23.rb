require 'pry'
require 'priority_queue'

class Bot
  attr_reader :x, :y, :z, :r

  def initialize(x, y, z, r)
    @x = Integer(x)
    @y = Integer(y)
    @z = Integer(z)
    @r = Integer(r)
  end

  def in_range?(bot)
    dist = (x - bot.x).abs + (y - bot.y).abs + (z - bot.z).abs

    dist <= r
  end
end

def parse(input)
  bots = []
  input.split("\n").each do |l|
    x, y, z, r = l.scan(/-?\d+/)
    bots << Bot.new(x, y, z, r)
  end

  bots
end

def run(input)
  bots = parse(input)

  max_bot = bots.max_by(&:r)

  in_range = bots.select do |bot|
               max_bot.in_range?(bot)
             end

  in_range.size
end

def run2(input)
  bots = parse(input)

  q = PriorityQueue.new

  bots.each do |bot|
    dist = bot.x.abs + bot.y.abs + bot.z.abs

    key1 = [1, bot.x, bot.y, bot.z]
    key2 = [-1, bot.x, bot.y, bot.z]
    q.push(key1, ([0, dist - bot.r].max))
    q.push(key2, (dist + bot.r + 1))
  end

  count = 0
  max = 0
  res = 0

  while !q.empty?
    key, dist = q.delete_min

    count += key[0]

    if count > max
      res = dist
      max = count
    end
  end

  res
end

def test
  input =<<~INPUT
pos=<0,0,0>, r=4
pos=<1,0,0>, r=1
pos=<4,0,0>, r=3
pos=<0,2,0>, r=1
pos=<0,5,0>, r=3
pos=<0,0,3>, r=1
pos=<1,1,1>, r=1
pos=<1,1,2>, r=1
pos=<1,3,1>, r=1
INPUT

  res = run(input)
  puts "1: EXPECTED #{res} to eq 7"
end

def test2
  input = <<~INPUT
pos=<10,12,12>, r=2
pos=<12,14,12>, r=2
pos=<16,12,12>, r=4
pos=<14,14,14>, r=6
pos=<50,50,50>, r=200
pos=<10,10,10>, r=5
INPUT
  res = run2(input.chomp)
  puts "2: EXPECTED #{res} to eq 36"
end


test
test2

input = File.read('23.input').chomp

puts "PART ONE: #{run(input)}"
puts "PART TWO: #{run2(input)}"
