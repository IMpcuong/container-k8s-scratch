class Coordinates {
  double latitude
  double longitude

  // NOTE: if you remove this fucntion, this error will be appearred.
  // -> `groovy.lang.MissingMethodException: No signature of method: Coordinates.getAt() is applicable for argument types: (Integer) values: [0]`.
  double getAt(int idx) {
    if (idx == 0) latitude
    else if (idx == 1) longitude
    else throw new Exception("Wrong coordinate index, use 0 or 1")
  }
}

def coordinates = new Coordinates(latitude: 43.23, longitude: 3.67)
def newObj = new Object();

// NOTE: index based object destructuring in Groovy.
def (la, lo) = coordinates

assert la == 43.23
assert lo == 3.67

void getObjProps(obj) {
  println obj.properties
    .sort{it.key}
    .collect{it}
    .findAll{!['class', 'active'].contains(it.key)}
    .join('\n')
}

getObjProps(coordinates)
// Output:
// latitude=43.23
// longitude=3.67

getObjProps(newObj)

// Run commands with piping operand: (resource ~ https://gist.github.com/lyuboraykov/8deae849e4812669793a)
def cmd = '''
docker ps -a | tail -n+2 | head -n4 | awk '{print $1}' | xargs docker container rm -f
'''
// def proc = ['bash', '-c', cmd].execute()
// println proc.text

def cmd1 = '''
docker rmi -f `docker images -a | grep -i pipenv | awk '{print $3}'`
'''
// def proc1 = ['bash', '-c', cmd1].execute()
// println proc1.text

def cmd2 = '''
docker images -a | grep -i qualys
docker run -d --restart=always qualys/sensor:1.19.0
'''
// def proc2 = ['bash', '-c', cmd2].execute()
// println proc2.text

def runPipeCmd(cmd) {
  def proc = ['bash', '-c', cmd].execute()
  return proc.text
}

println 'docker images -a'.execute().text
println 'docker ps -a'.execute().text

def rmCmd = '''
docker images -a | tail -n+2 | awk '{if($1 ~ /^.*none/ || $1 ~ /^.*null/) print $3}' | xargs docker rmi -f
'''
runPipeCmd(rmCmd)
