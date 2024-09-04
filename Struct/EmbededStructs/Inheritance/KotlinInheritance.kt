import kotlinx.serialization.Serializable
import kotlinx.serialization.encodeToString
import kotlinx.serialization.json.Json

@Serializable
open class Product(
    var name: String,
    var color: String,
    var length: Int,
    var width: Int,
    var weight: Int,
    var price: Int,
    var brand: String,
    var madeIn: String
)

@Serializable
open class ElectronicProduct(
    name: String,
    color: String,
    length: Int,
    width: Int,
    weight: Int,
    price: Int,
    brand: String,
    madeIn: String,
    var ram: Int,
    var cpu: String,
    var screenSize: Float,
    var operatingSystem: String,
    var operatingSystemVersion: String
) : Product(name, color, length, width, weight, price, brand, madeIn)

@Serializable
class Mobile(
    name: String,
    color: String,
    length: Int,
    width: Int,
    weight: Int,
    price: Int,
    brand: String,
    madeIn: String,
    ram: Int,
    cpu: String,
    screenSize: Float,
    operatingSystem: String,
    operatingSystemVersion: String,
    var simcardCapacity: Int,
    var simcardType: String,
    var networkType: String,
    var cameraCount: Int
) : ElectronicProduct(name, color, length, width, weight, price, brand, madeIn, ram, cpu, screenSize, operatingSystem, operatingSystemVersion)

@Serializable
class Loptop(
    name: String,
    color: String,
    length: Int,
    width: Int,
    weight: Int,
    price: Int,
    brand: String,
    madeIn: String,
    ram: Int,
    cpu: String,
    screenSize: Float,
    operatingSystem: String,
    operatingSystemVersion: String,
    var usbPortCount: Int,
    var keyboardType: String,
    var hasCdRom: Boolean
) : ElectronicProduct(name, color, length, width, weight, price, brand, madeIn, ram, cpu, screenSize, operatingSystem, operatingSystemVersion)

fun main() {
    val mobile = Mobile(
        name = "Samsung M4150",
        color = "Blue",
        length = 0,
        width = 0,
        weight = 0,
        price = 0,
        brand = "Samsung",
        madeIn = "",
        ram = 4,
        cpu = "",
        screenSize = 0f,
        operatingSystem = "Android",
        operatingSystemVersion = "11",
        simcardCapacity = 3,
        simcardType = "Nano",
        networkType = "5G",
        cameraCount = 2
    )

    val loptop = Loptop(
        name = "MSI GX 4570",
        color = "Black",
        length = 0,
        width = 0,
        weight = 0,
        price = 0,
        brand = "MSI",
        madeIn = "",
        ram = 16,
        cpu = "",
        screenSize = 0f,
        operatingSystem = "Ubuntu",
        operatingSystemVersion = "22.04",
        usbPortCount = 3,
        keyboardType = "Light",
        hasCdRom = false
    )

    val mobileJson = Json.encodeToString(mobile)
    val loptopJson = Json.encodeToString(loptop)

    println(mobileJson)
    println(loptopJson)
}
