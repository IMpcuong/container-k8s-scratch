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