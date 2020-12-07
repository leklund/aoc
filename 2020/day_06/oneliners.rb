s = File.read("input.txt")
puts s.split("\n\n").map{|g| g.split(/\n?/).uniq.count}.sum
puts s.split("\n\n").map{|g| g.split("\n").map{|z| z.split("")}.inject(:'&').count}.sum
