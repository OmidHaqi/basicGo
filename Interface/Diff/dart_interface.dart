abstract class IPerson {
  String name;
  int age;

  void printInfo() {
    print('$name $age');
  }
}

class Person implements IPerson {
  @override
  String name;

  @override
  int age;

  Person(this.name, this.age);

  @override
  void printInfo() {
    print('$name - $age');
  }
}

void main() {
  Person person = Person('John Doe', 25);
  person.printInfo();
}
