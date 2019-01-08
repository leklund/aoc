require 'pry'

# instead of building a binary serach and adding loop detection for infinte loops
# I just did a manual binary search by hand.
BOOST = 52

class Army
  attr_reader :id, :units, :weaknesses, :immunities, :hp, :damage, :attack, :type, :initiative
  attr_accessor :units, :current_target, :targeted, :boost

  def initialize(units:, hp:, damage:, weaknesses: nil, immunities: nil, initiative:, attack:, type:, id:)
    @units = units
    @hp = hp
    @damage = damage
    @weaknesses = weaknesses
    @immunities = immunities
    @initiative = initiative
    @attack = attack
    @type = type
    @targeted = false
    @id = id
    @boost = type == :immune ? BOOST : 0
  end

  def effective_power
    units * (damage + boost)
  end

  def clear
    @targeted = false
    @current_target = nil
  end

  def attack_damage(enemy)
    if enemy.immunities&.include?(attack)
      0
    elsif enemy.weaknesses&.include?(attack)
      effective_power * 2
    else
      effective_power
    end
  end

  def target(enemies = nil)
    return current_target if enemies.nil?

    # damage
    targs = enemies.each_with_object(Hash.new { |h, k| h[k] = Array.new }) do |enemy, h|
      d = if enemy.targeted
            0
          else
            attack_damage(enemy)
          end

      h[d] << enemy
    end

    max = targs.keys.max
    possible_targets = targs[max]

    if possible_targets.empty? || max == 0
      self.current_target = nil
      return current_target
    elsif possible_targets.size == 1
      self.current_target = possible_targets.first
      current_target.targeted = true
      return current_target
    end

    # TIEBREAKERS
    new_targ = possible_targets.min_by { |t| [-t.effective_power, -t.initiative] }

    self.current_target = new_targ
    current_target.targeted = true
  end

  def attack!
    return unless current_target
    return if current_target.units.zero?

    dmg = attack_damage(current_target)
    z = current_target&.units
    # puts "#{type} #{id} - attacking: #{current_target&.type} #{current_target&.id} | #{current_target&.units} with #{dmg} damage"
    current_target.take_damage(dmg)
    # puts "      killed #{z - current_target&.units} units"
  end

  def take_damage(hit_points)
    
    remaining = hp * units - hit_points
    if remaining <= 0
      @units = 0
    else
      @units = (remaining / hp.to_f).ceil
    end
  end
end

def run(input)
  input.chomp
  armies = input.split("\n\n")
  immune = parse(armies.first, :immune)
  infection = parse(armies.last, :infection)

  armies = immune + infection

  battle(armies)
end

def battle(armies)
  battle_on = true

  # DEBUG
  # puts "get ready to rumble"
  # immune = armies.select { |a| a.type == :immune }
  # infection = armies.select { |a| a.type == :infection}
  # puts "IMMUNE"
  # immune.each { |a| puts "#{a.id} - #{a.units} units, #{a.hp} hp, #{a.weaknesses} weak, #{a.immunities} immune, #{a.damage} damage, #{a.attack} attack" }
  # puts "INFECTION"
  # infection.each { |a| puts "#{a.id} - #{a.units} units, #{a.hp} hp, #{a.weaknesses} weak, #{a.immunities} immune, #{a.damage} damage, #{a.attack} attack" }

  while battle_on
    armies.sort_by! { |a| [-a.effective_power, -a.initiative] }

    # TARGETING PHASE
    immune = armies.select { |a| a.type == :immune }
    infection = armies.select { |a| a.type == :infection }

    armies.each do |army|
      enemies = army.type == :immune ? infection : immune
      army.target(enemies)
      # puts "#{army.type} #{army.id} - u: #{army.units} ep: #{army.effective_power}, i: #{army.initiative}, target: #{army.current_target&.type} | #{army.current_target&.units}"
    end

    # ATTACK
    armies.sort_by(&:initiative).reverse_each do |army|
      army.attack! if army.units >= 0
    end

    armies.select! { |a| a.units > 0 }

    if armies.map(&:type).uniq.size == 1
      # winner
      return armies
    end

    armies.each(&:clear)
    # puts "----------------------------"
  end
end

def parse(lines, type)
  lines = lines.split("\n")
  lines.shift
  lines.each_with_object([]).with_index do |(line, army), i|
    matches = line.match(/^(\d+) units each with (\d+) hit points (?:\((?:(immune|weak) to (.+?)(?:;|\)) )?(weak|immune) to (.+)\))?.+does (\d+) (\w+) damage at initiative (\d+)/)
    # they just had to be in any order didn't they
    imm_idx = 4 if matches[3] == 'immune'
    imm_idx = 6 if matches[5] == 'immune'
    weak_idx = 4 if matches[3] == 'weak'
    weak_idx = 6 if matches[5] == 'weak'
    immunities = matches[imm_idx]&.split(', ')&.map(&:to_sym) if imm_idx
    weaks = matches[weak_idx]&.split(', ')&.map(&:to_sym) if weak_idx
    army << Army.new(units: matches[1].to_i,
                     hp: matches[2].to_i,
                     immunities: immunities,
                     weaknesses: weaks,
                     damage: matches[7].to_i,
                     attack: matches[8].to_sym,
                     initiative: matches[9].to_i,
                     type: type,
                     id: i + 1,)
  end
end

def test
  input =<<~INPUT
Immune System:
17 units each with 5390 hit points (weak to radiation, bludgeoning) with an attack that does 4507 fire damage at initiative 2
989 units each with 1274 hit points (immune to fire; weak to bludgeoning, slashing) with an attack that does 25 slashing damage at initiative 3

Infection:
801 units each with 4706 hit points (weak to radiation) with an attack that does 116 bludgeoning damage at initiative 1
4485 units each with 2961 hit points (immune to radiation; weak to fire, cold) with an attack that does 12 slashing damage at initiative 4
INPUT

  winners = run(input)

  puts "expected #{winners.map(&:units).sum} to eq 5216"
end

test

input = DATA.read

# SET BOOST = 0 to run part 1
winners = run(input)
puts "WINNERS: #{winners.first.type}"
puts "UNITS LEFT: #{winners.map(&:units).sum}"

__END__
Immune System:
2208 units each with 6238 hit points (immune to slashing) with an attack that does 23 bludgeoning damage at initiative 20
7603 units each with 6395 hit points (weak to radiation) with an attack that does 6 cold damage at initiative 15
4859 units each with 5904 hit points (weak to fire) with an attack that does 12 cold damage at initiative 11
1608 units each with 7045 hit points (weak to fire, cold; immune to bludgeoning, radiation) with an attack that does 31 radiation damage at initiative 10
39 units each with 4208 hit points with an attack that does 903 radiation damage at initiative 7
6969 units each with 9562 hit points (immune to slashing, cold) with an attack that does 13 slashing damage at initiative 3
2483 units each with 6054 hit points (immune to fire) with an attack that does 20 cold damage at initiative 19
506 units each with 3336 hit points with an attack that does 64 radiation damage at initiative 6
2260 units each with 10174 hit points (weak to fire) with an attack that does 34 slashing damage at initiative 5
2817 units each with 9549 hit points (immune to cold, fire; weak to bludgeoning) with an attack that does 31 cold damage at initiative 2

Infection:
3650 units each with 25061 hit points (weak to fire, bludgeoning) with an attack that does 11 slashing damage at initiative 12
508 units each with 48731 hit points (weak to bludgeoning) with an attack that does 172 cold damage at initiative 13
724 units each with 27385 hit points with an attack that does 69 radiation damage at initiative 1
188 units each with 41786 hit points with an attack that does 416 bludgeoning damage at initiative 4
3045 units each with 36947 hit points (weak to slashing; immune to fire, bludgeoning) with an attack that does 24 slashing damage at initiative 9
7006 units each with 42545 hit points (immune to cold, slashing, fire) with an attack that does 9 fire damage at initiative 16
853 units each with 55723 hit points (weak to cold, fire) with an attack that does 114 bludgeoning damage at initiative 17
3268 units each with 43027 hit points (immune to slashing, fire) with an attack that does 25 slashing damage at initiative 8
1630 units each with 47273 hit points (weak to cold, bludgeoning) with an attack that does 57 slashing damage at initiative 14
3383 units each with 12238 hit points with an attack that does 7 radiation damage at initiative 18
