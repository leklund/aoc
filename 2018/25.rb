class UFind
  attr_reader :parents, :sizes
  attr_accessor :sets

  def initialize(arr)
    @parents = arr.each_with_object({}) { |e, h| h[e] = e }
    @sizes = Hash.new(1)
    @sets = 0
  end

  def find_root(node)
    return parents[node] if parents[node] == node

    # path compresson
    parents[node] = find_root(parents[node])
  end

  def union(x, y)
    xroot = find_root(x)
    yroot = find_root(y)

    return if xroot == yroot

    if sizes[x] < sizes[y]
      # swap xroot yroot
      xroot, yroot = yroot, xroot
    end

    # union
    parents[yroot] = xroot
    sizes[xroot] += sizes[yroot]
  end
end

def parser(input)
  input.split("\n").map do |line|
    line.scan(/-?\d+/).map(&:to_i)
  end
end

def run(input)
  input.chomp

  galaxy = parser(input)

  union_find = UFind.new(galaxy)

  # check every pair of stars and union if they are within 3 manhattan distance
  galaxy.combination(2).each do |star_x, star_y|
    md = star_x.zip(star_y).sum { |a, b| (a - b).abs }
    union_find.union(star_x, star_y) if md <= 3
  end

  # I thought this would work but I'm not sure why it doesn't
  # union_find.parents.values.uniq.size
  #
  # seems we have to find root on every star to do some path compresion and now I can
  # count the parents. Could also count if root == star
  galaxy.each { |star| union_find.find_root(star) }
  union_find.parents.values.uniq.size
end

def test
  input1 =<<~INPUT
0,0,0,0
3,0,0,0
0,3,0,0
0,0,3,0
0,0,0,3
0,0,0,6
9,0,0,0
12,0,0,0
INPUT

  input2 =<<~INPUT
-1,2,2,0
0,0,2,-2
0,0,0,-2
-1,2,0,0
-2,-2,-2,2
3,0,2,-1
-1,3,2,2
-1,0,-1,0
0,2,1,-2
3,0,0,0
INPUT

  input3 =<<~INPUT
1,-1,0,1
2,0,-1,0
3,2,-1,0
0,0,3,1
0,0,-1,-1
2,3,-2,0
-2,2,0,0
2,-2,0,-1
1,-1,0,-1
3,2,0,2
INPUT

  input4 =<<~INPUT
1,-1,-1,-2
-2,-2,0,1
0,2,1,3
-2,3,-2,1
0,2,3,-2
-1,-1,1,-2
0,-2,-1,0
-2,2,3,-1
1,2,2,0
-1,-2,0,-2
  INPUT

  res = run(input1)
  puts "EXPECT #{res} to eq 2"

  res = run(input2)
  puts "EXPECT #{res} to eq 4"

  res = run(input3)
  puts "EXPECT #{res} to eq 3"

  res = run(input4)
  puts "EXPECT #{res} to eq 8"

end

test

input = File.read('25.input')

puts run(input)
