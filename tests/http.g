std.write "#######################################################" end
std.write "" "" end
std.write "Testing gecko.http package..." end
std.write "" "" end
std.write "#######################################################" end
std.write "" "" end

#
#
# test data
#
#

def "get_response" value
  map
    "userId" 1
    "id" 1
    "title" "delectus aut autem"
    "completed" false
  end
end

def "post_response" value
  map
    "id" 101
  end
end

def "test_value" value nil end

#
#
# get request
#
#

let "test_value" value
  http
    "https://jsonplaceholder.typicode.com/todos/1"
    "GET"
    map
      "Accept" "application/json"
    end
    map end
  end
end

def "get_worked" value
  .map.equal
    test_value
    get_response
  end
end

if
  get_worked
then
  std.write "PASS: http get request" end
else
  std.write "FAIL: http get request" end
end

#
#
# post request
#
#

let "test_value" value
  http
    "https://jsonplaceholder.typicode.com/posts"
    "POST"
    map
      "Accept" "application/json"
    end
    map
      "title" "foo"
      "body" "bar"
      "userId" 1
    end
  end
end

def "post_worked" value
  .map.equal
    test_value
    post_response
  end
end

if
  post_worked
then
  std.write "PASS: http post request" end
else
  std.write "FAIL: http post request" end
end