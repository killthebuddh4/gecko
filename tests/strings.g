std.write "#######################################################" end
std.write "" end
std.write "" end
std.write "Testing gecko.strings package..." end
std.write "" end
std.write "" end
std.write "#######################################################" end
std.write "" end
std.write "" end

#
#
# test data
#
#

def "lang" value "gecko" end

def "splitted" value
  array "g" "a" "d" "f" "l" "y" end
end

def "test_value" value nil end

#
#
# strings.Split
#
#

let "test_value" value
  split lang end
end

def "split_worked" value
  .equal test_value splitted end
end

if
  split_worked == false
then
  std.write "FAIL: strings.Split" end
else
  std.write "PASS: strings.Split" end
end

#
#
# strings.Substring
#
#

let "test_value" value
  substring lang 3 6 end
end

if
  test_value == "fly"
then
  std.write "PASS: strings.Substring" end
else
  std.write "FAIL: strings.Substring" end
end

#
#
# strings.Concat
#
#

def "gad" value "gad" end

let "test_value" value
  concat
    gad
    "fly"
  end
end

if
  test_value == "gecko"
then
  std.write "PASS: strings.Concat" end
else
  std.write "FAIL: strings.Concat" end
end
