# frozen_string_literal: true

require 'pry'

class Player
  attr_accessor :score
  def initialize
    @score = 0
  end
end

def next_index(arr, i)
  return 0 if i.nil? # first loop
  return 1 if arr[i + 1].nil? # at tail?

  i + 2
end

def run(input)
  player_count, last_val = input.scan(/\d+/)

  players = player_count.to_i.times.map { Player.new }
  players_loop = players.cycle
  arr = []
  curr_index = nil
  (0..last_val.to_i).each do |val|
    active_player = players_loop.next

    if !val.zero? && val % 23 == 0
      # SCORE!
      active_player.score += val
      curr_index = if curr_index < 7
                     arr.size - curr_index - 1
                   else
                     curr_index -= 7
                   end

      active_player.score += arr.delete_at(curr_index)

      curr_index = 0 if arr.size == curr_index # did I delete the last element?
    else
      curr_index = next_index(arr, curr_index)

      arr.insert(curr_index, val)
    end
  end
  players.max_by(&:score).score
end

def test(input, expected)
  actual = run(input)

  puts "❌ expected #{actual} to eq #{expected}" unless actual == expected

  #puts '✅'
end

test('9 players; 25 points', 32)
test('10 players; last marble is worth 1618 points', 8317)
test('13 players; last marble is worth 7999 points', 146373)
test('17 players; last marble is worth 1104 points', 2764)
test('21 players; last marble is worth 6111 points', 54718)
test('30 players; last marble is worth 5807 points', 37305)

puts run ('412 players; last marble is worth 71646 points')
puts run ('412 players; last marble is worth 7164600 points') 
