# I started building a LinkedList for part one and then decided to
# brute force it. This implementation also means that all the test
# cases from part one pass (the second example didn't pass with my array
# implementation).
#
# This might have been better with a collection
# of node objects and cursor to track current postition
class Linky
  attr_accessor :head, :tail
  attr_reader :data
  
  def self.init
    root = new(0)
    root.head = root
    root.tail = root
    root
  end

  def initialize(data, head = nil, tail = nil)
    @data = data
    @head = head
    @tail = tail
  end

  def add(data)
    node = self.class.new(data, self, tail)
    self.head = node if head == tail

    tail.head = node
    self.tail = node

    node
  end

  def delete
    node = self
    node.head.tail = node.tail
    node.tail.head = node.head
    node.tail
  end

  def rewind(x)
    node = self
    x.times { node = node.head }
    node
  end
end

class Player
  attr_accessor :score
  def initialize
    @score = 0
  end
end

def run(input)
  player_count, last_val = input.scan(/\d+/)

  players = player_count.to_i.times.map { Player.new }
  players_loop = players.cycle

  list = Linky.init
  (1..last_val.to_i).each do |val|
    active_player = players_loop.next

    if val % 23 == 0
      # SCORE!
      active_player.score += val
      list = list.rewind(7)

      active_player.score += list.data

      list = list.delete
    else
      list = list.tail
      list = list.add(val)
    end
  end

  players.max_by(&:score).score
end

def test(input, expected)
  actual = run(input)

  puts "❌ expected #{actual} to eq #{expected}" unless actual == expected

  puts '✅' if actual == expected
end

test('9 players; 25 points', 32)
test('10 players; last marble is worth 1618 points', 8317)
test('13 players; last marble is worth 7999 points', 146373)
test('17 players; last marble is worth 1104 points', 2764)
test('21 players; last marble is worth 6111 points', 54718)
test('30 players; last marble is worth 5807 points', 37305)

puts run ('412 players; last marble is worth 71646 points')
puts run ('412 players; last marble is worth 7164600 points') 
