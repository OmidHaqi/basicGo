interface IPerson {
    var name: String
    var age: Int

    fun printInfo() {
        println("$name $age")
    }
}

class Person(override var name: String, override var age: Int) : IPerson {
    override fun printInfo() {
        println("$name - $age")
    }
}

fun main() {
    val person = Person("John Doe", 25)
    person.printInfo()
}
