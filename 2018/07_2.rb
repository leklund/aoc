class Node
  attr_reader :id
  attr_accessor :parents, :children, :completed, :in_progress

  OFFSET = 4

  def initialize(id:, parents: [], children: [])
    @id = id
    @parents = parents
    @children = children
    @completed = false
    @in_progress = false
  end

  def complete
    self.completed = true
    self.in_progress = false
  end

  def complete?
    @completed
  end

  def incomplete?
    !complete?
  end

  def to_i
    self.id.ord - OFFSET
  end
end

class Tree
  attr_accessor :nodes

  def initialize
    @nodes ||= []
  end

  def [](id)
    node = nodes.find { |n| n.id == id }
    if node.nil?
      node = Node.new(id: id)
      @nodes << node
    end

    node
  end

  def ready
    nodes.select { |node| node.parents.empty? && node.incomplete? && !node.in_progress }.sort_by(&:id)
  end

  def remove_dependency(dep)
    nodes.select { |node| node.parents.include?(dep.id) }.each { |node| node.parents.delete(dep.id) }
  end
end

class Elf
  attr_accessor :working, :ttl, :current_step

  def initialize
    @working = false
  end

  def ready?
    self.working == false
  end

  def working?
    !ready?
  end
end

def run(input)
  tree = Tree.new
  input.each do |line|
    parent, child = line.scan(/Step (\w).*before step (\w)/).flatten

    tree[child].parents << parent
    tree[parent].children << child
  end

  end_node = tree.nodes.find { |n| n.children.empty? }
  t = 0
  workers = 5.times.map { Elf.new }

  until end_node.complete?
    nodes = tree.ready

    # assign workers
    workers.select(&:ready?).each do |elf|
      next unless node = nodes.shift
      node.in_progress = true
      elf.current_step = node
      elf.ttl = node.to_i
      elf.working = true
    end

    curr = workers.map { |w| w.current_step&.id || 0 }
    puts "#{t}: #{curr.join(' ')}"

    # do work
    t += 1
    workers.select(&:working?).each do |elf|
      elf.ttl -= 1
      if elf.ttl == 0
        elf.current_step.complete
        tree.remove_dependency(elf.current_step)
        elf.current_step = nil
        elf.working = false
      end
    end
  end

  t
end

def test
  input = <<~INPUT
    Step C must be finished before step A can begin.
    Step C must be finished before step F can begin.
    Step A must be finished before step B can begin.
    Step A must be finished before step D can begin.
    Step B must be finished before step E can begin.
    Step D must be finished before step E can begin.
    Step F must be finished before step E can begin.
  INPUT

  expected = 15
  actual = run(input.split("\n"))

  abort "❌ expected #{actual} to eq #{expected}" unless actual == expected

  puts '✅'
end

# Can't run test because the OFFSET is different for real data
# test

data = DATA.read.chomp.split("\n")

out = run(data)
puts out

__END__
Step G must be finished before step W can begin.
Step X must be finished before step S can begin.
Step F must be finished before step V can begin.
Step C must be finished before step Y can begin.
Step M must be finished before step J can begin.
Step K must be finished before step Z can begin.
Step U must be finished before step W can begin.
Step I must be finished before step H can begin.
Step W must be finished before step B can begin.
Step A must be finished before step Y can begin.
Step Y must be finished before step D can begin.
Step S must be finished before step Q can begin.
Step N must be finished before step V can begin.
Step H must be finished before step D can begin.
Step D must be finished before step Q can begin.
Step L must be finished before step E can begin.
Step Q must be finished before step E can begin.
Step T must be finished before step R can begin.
Step J must be finished before step P can begin.
Step R must be finished before step E can begin.
Step E must be finished before step V can begin.
Step O must be finished before step P can begin.
Step P must be finished before step B can begin.
Step Z must be finished before step V can begin.
Step B must be finished before step V can begin.
Step Y must be finished before step B can begin.
Step C must be finished before step B can begin.
Step Q must be finished before step T can begin.
Step W must be finished before step P can begin.
Step X must be finished before step Z can begin.
Step L must be finished before step T can begin.
Step G must be finished before step Y can begin.
Step Y must be finished before step R can begin.
Step E must be finished before step B can begin.
Step X must be finished before step E can begin.
Step Y must be finished before step V can begin.
Step H must be finished before step L can begin.
Step L must be finished before step J can begin.
Step S must be finished before step T can begin.
Step F must be finished before step T can begin.
Step Y must be finished before step J can begin.
Step A must be finished before step H can begin.
Step P must be finished before step Z can begin.
Step R must be finished before step O can begin.
Step X must be finished before step F can begin.
Step I must be finished before step O can begin.
Step Y must be finished before step Q can begin.
Step S must be finished before step D can begin.
Step Q must be finished before step B can begin.
Step C must be finished before step D can begin.
Step Y must be finished before step N can begin.
Step O must be finished before step Z can begin.
Step G must be finished before step D can begin.
Step A must be finished before step O can begin.
Step U must be finished before step N can begin.
Step Y must be finished before step P can begin.
Step E must be finished before step O can begin.
Step I must be finished before step Q can begin.
Step W must be finished before step O can begin.
Step D must be finished before step B can begin.
Step Z must be finished before step B can begin.
Step L must be finished before step B can begin.
Step P must be finished before step V can begin.
Step C must be finished before step E can begin.
Step S must be finished before step O can begin.
Step U must be finished before step T can begin.
Step U must be finished before step O can begin.
Step Y must be finished before step L can begin.
Step N must be finished before step L can begin.
Step Q must be finished before step Z can begin.
Step U must be finished before step L can begin.
Step U must be finished before step D can begin.
Step J must be finished before step O can begin.
Step L must be finished before step R can begin.
Step S must be finished before step P can begin.
Step H must be finished before step R can begin.
Step F must be finished before step I can begin.
Step D must be finished before step T can begin.
Step C must be finished before step M can begin.
Step W must be finished before step D can begin.
Step R must be finished before step V can begin.
Step U must be finished before step S can begin.
Step K must be finished before step R can begin.
Step D must be finished before step V can begin.
Step D must be finished before step R can begin.
Step I must be finished before step E can begin.
Step L must be finished before step O can begin.
Step T must be finished before step Z can begin.
Step A must be finished before step E can begin.
Step D must be finished before step Z can begin.
Step H must be finished before step V can begin.
Step A must be finished before step L can begin.
Step W must be finished before step R can begin.
Step F must be finished before step A can begin.
Step Y must be finished before step Z can begin.
Step I must be finished before step P can begin.
Step F must be finished before step J can begin.
Step H must be finished before step B can begin.
Step G must be finished before step Z can begin.
Step C must be finished before step K can begin.
Step D must be finished before step E can begin.
