std.write "#######################################################" end
std.write "" "" end
std.write "Testing gecko.array package..." end
std.write "" "" end
std.write "#######################################################" end
std.write "" "" end

#
#
# test data
#
#

def "numbers" value
  array
    2 4 6 7 8 10
  end
end


def "numbers_clone" value
  array.map
    numbers

    (ab: Number i)
      fn
        ab
      end
  end
end

def "numbers_and_twenty" value
  array 2 4 6 7 8 10 20 end
end

def "numbers_without_ten" value
  array 2 4 6 7 8 end
end

def "numbers_without_two" value
  array 4 6 7 8 10 end
end

def "numbers_with_zero" value
  array 0 2 4 6 7 8 10 end
end

def "numbers_sum" value 37 end

def "squares" value
  array 4 16 36 49 64 100 end
end

def "evens" value
  array 2 4 6 8 10 end
end

def "test_val" value nil end

#
#
# array.Read
#
#

let "test_val" value
  array.read numbers 2 end
end

std.write
  if
    test_val != 6
  then
    "FAIL: array.Read"
  else
    "PASS: array.Read"
  end
end

#
#
# array.Write
#
#

array.write numbers 2 12 end

let "test_val" value
  array.read numbers 2 end
end

std.write
  if
    test_val != 6
  then
    "FAIL: array.Write immutability"
  else
    "PASS: array.Write immutability"
  end
end

def "new_numbers" value
  array.write numbers 2 12 end
end

let "test_val" value
  array.read new_numbers 2 end
end

if
  test_val != 12
then
  std.write "FAIL: array.Write" end
else
  std.write "PASS: array.Write" end
end

#
#
# array.For
#
#

let "test_val" value 0 end

array.for numbers
  (a i)
    fn
      let "test_val" value test_val + a end
    end
end

if
  test_val != numbers_sum
then
  std.write "FAIL: array.For" end
else
  std.write "PASS: array.For" end
end

#
#
# array.Map
#
#

let "test_val" value
  array.map numbers
    (a i)
      fn
        a * a
      end
  end
end

def "squares_worked" value
  .equal test_val squares end
end

if
  squares_worked == false
then
  std.write "FAIL: array.Map" end
else
  std.write "PASS: array.Map" end
end 

#
#
# array.Filter
#
#

let "test_val" value
  array.filter numbers
    (a i)
      fn
        def "rem" value
          .math.mod a 2 end
        end

        rem == 0
      end
  end
end

def "evens_worked" value
  .equal test_val evens end
end

if
  evens_worked == false
then
  std.write "FAIL: array.Filter" end
else
  std.write "PASS: array.Filter" end
end

#
#
# array.Reduce
#
#

let "test_val" value
  array.reduce
    numbers

    0

    (acc a i)
      fn
        acc + a
      end
  end
end

if
  test_val != numbers_sum
then
  std.write "FAIL: array.Reduce" end
else
  std.write "PASS: array.Reduce" end
end

#
#
# array.Push
#
#

array.push numbers 20 end

def "push_is_immutable" value
  .equal numbers_clone numbers end
end

if
  push_is_immutable == false
then
  std.write "FAIL: array.Push immutability" end
else
  std.write "PASS: array.Push immutability" end
end

let "test_val" value
  array.push numbers 20 end
end

def "push_worked" value
  .equal test_val numbers_and_twenty end
end

if push_worked == false
then
  std.write "FAIL: array.Push" end
else
  std.write "PASS: array.Push" end
end

#
#
# array.Pop
#
#

array.pop numbers end

def "pop_is_immutable" value
  .equal numbers_clone numbers end
end

if pop_is_immutable == false
then
  std.write "FAIL: array.Pop immutability" end
else
  std.write "PASS: array.Pop immutability" end
end

let "test_val" value
  array.pop numbers end
end

def "pop_worked" value
  .equal test_val numbers_without_ten end
end

if pop_worked == false
then
  std.write "FAIL: array.Pop" end
else
  std.write "PASS: array.Pop" end
end

#
#
# array.Shift
#
#

array.shift numbers end

def "shift_is_immutable" value
  .equal numbers_clone numbers end
end

if
  shift_is_immutable == false
then
  std.write "FAIL: array.Shift immutability" end
else
  std.write "PASS: array.Shift immutability" end
end

let "test_val" value
  array.shift numbers end
end

def "shift_worked" value
  .equal test_val numbers_without_two end
end

if
  shift_worked == false
then
  std.write "FAIL: array.Shift" end
else
  std.write "PASS: array.Shift" end
end

#
#
# array.Unshift
#
#

array.unshift numbers 0 end

def "unshift_is_immutable" value
  .equal numbers_clone numbers end
end

if
  unshift_is_immutable == false
then
  std.write "FAIL: array.Unshift immutability" end
else
  std.write "PASS: array.Unshift immutability" end
end

let "test_val" value
  array.unshift numbers 0 end
end

def "unshift_worked" value
  .equal test_val numbers_with_zero end
end

if
  unshift_worked == false
then
  std.write "FAIL: array.Unshift" end
else
  std.write "PASS: array.Unshift" end
end