require 'scanf'

class Oper
  attr_reader :register, :instruction

  INSTRUCTIONS = %i(addr addi mulr muli banr bani borr bori setr seti gtir gtri gtrr eqir eqri eqrr)

  def initialize(instruction, register)
    @instruction = instruction
    @register = register
  end

  def self.method_missing(meth, *args)
    o = new(*args)
    o.send(meth)
    o.register
  end

  def code
    instruction[0]
  end

  def a
    instruction[1]
  end

  def b
    instruction[2]
  end

  def c
    instruction[3]
  end

  def addr
    register[c] = register[a] + register[b]
  end

  def addi
    register[c] = register[a] + b
  end

  def mulr
    register[c] = register[a] * register[b]
  end

  def muli
    register[c] = register[a] * b
  end

  def banr
    register[c] = register[a] & register[b]
  end

  def bani
    register[c] = register[a] & b
  end

  def borr
    register[c] = register[a] | register[b]
  end

  def bori
    register[c] = register[a] | b
  end

  def setr
    register[c] = register[a]
  end

  def seti
    register[c] = a
  end

  def gtir
    register[c] = a > register[b] ? 1 : 0
  end

  def gtri
    register[c] = register[a] > b ? 1 : 0
  end

  def gtrr
    register[c] = register[a] > register[b] ? 1 : 0
  end

  def eqir
    register[c] = a == register[b] ? 1 : 0
  end

  def eqri
    register[c] = register[a] == b ? 1 : 0
  end

  def eqrr
    register[c] = register[a] == register[b] ? 1 : 0
  end
end

def parse(input)
  return if input[0].nil?
  before = input[0].scanf('Before: [%d, %d, %d, %d]')
  after = input[2].scanf('After:  [%d, %d, %d, %d]')
  instruction = input[1].scanf('%d %d %d %d')
  [before, after, instruction]
end

def calculate_opcodes(ophash)
  codes = {}
  opcodes = Oper::INSTRUCTIONS.clone
  until opcodes.size == 0
    uniq = ophash.sort_by { |_k, v| v.size }.first
    id = uniq.first
    op = uniq.last.keys.first

    codes[id] = op
    opcodes.delete(op)
    ophash.delete(id)
    ophash.values.each { |vh| vh.delete(op) }
  end

  codes
end

def run(input)
  samples = 0

  ops = Hash.new { |h, k| h[k] = Hash.new(0) }

  loop do
    break if input[0].nil? || input[0].empty?
    before, after, instruction, _l = parse(input.shift(4))

    break if before.nil? || before.empty?

    count = 0
    Oper::INSTRUCTIONS.each do |op|
      if after == Oper.send(op, instruction, before.clone)
        ops[instruction[0]][op] += 1
        count += 1
      end
    end
    samples += 1 if count >= 3
  end

  puts "samples > 3 ops: #{samples}"

  codemap = calculate_opcodes(ops)

  # clear the blank lines
  while input[0] == ''
    input.shift
  end

  register = Array.new(4, 0)
  # run the program
  input.each do |line|
    instruction = line.scanf('%d %d %d %d')
    op = codemap[instruction[0]]
    Oper.send(op, instruction, register)
  end

  register
end

def test
  input = <<~INPUT
Before: [3, 2, 1, 1]
9 2 1 2
After:  [3, 2, 2, 1]

Before: [3, 2, 1, 1]
9 2 1 2
After:  [3, 2, 2, 1]


  INPUT
  input = input.split("\n")
  run(input)
end

# test

data = File.read('16_input.rb')
data = data.chomp.split("\n")
result = run(data)

puts result.inspect
