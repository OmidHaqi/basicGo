import 'dart:convert';

class Product {
  String name;
  String color;
  int length;
  int width;
  int weight;
  int price;
  String brand;
  String madeIn;

  Product({
    required this.name,
    required this.color,
    required this.length,
    required this.width,
    required this.weight,
    required this.price,
    required this.brand,
    required this.madeIn,
  });

  Map<String, dynamic> toJson() => {
        'name': name,
        'color': color,
        'length': length,
        'width': width,
        'weight': weight,
        'price': price,
        'brand': brand,
        'madeIn': madeIn,
      };
}

class ElectronicProduct extends Product {
  int ram;
  String cpu;
  double screenSize;
  String operatingSystem;
  String operatingSystemVersion;

  ElectronicProduct({
    required String name,
    required String color,
    required int length,
    required int width,
    required int weight,
    required int price,
    required String brand,
    required String madeIn,
    required this.ram,
    required this.cpu,
    required this.screenSize,
    required this.operatingSystem,
    required this.operatingSystemVersion,
  }) : super(
          name: name,
          color: color,
          length: length,
          width: width,
          weight: weight,
          price: price,
          brand: brand,
          madeIn: madeIn,
        );

  @override
  Map<String, dynamic> toJson() => super.toJson()
    ..addAll({
      'ram': ram,
      'cpu': cpu,
      'screenSize': screenSize,
      'operatingSystem': operatingSystem,
      'operatingSystemVersion': operatingSystemVersion,
    });
}

class Mobile extends ElectronicProduct {
  int simcardCapacity;
  String simcardType;
  String networkType;
  int cameraCount;

  Mobile({
    required String name,
    required String color,
    required int length,
    required int width,
    required int weight,
    required int price,
    required String brand,
    required String madeIn,
    required int ram,
    required String cpu,
    required double screenSize,
    required String operatingSystem,
    required String operatingSystemVersion,
    required this.simcardCapacity,
    required this.simcardType,
    required this.networkType,
    required this.cameraCount,
  }) : super(
          name: name,
          color: color,
          length: length,
          width: width,
          weight: weight,
          price: price,
          brand: brand,
          madeIn: madeIn,
          ram: ram,
          cpu: cpu,
          screenSize: screenSize,
          operatingSystem: operatingSystem,
          operatingSystemVersion: operatingSystemVersion,
        );

  @override
  Map<String, dynamic> toJson() => super.toJson()
    ..addAll({
      'simcardCapacity': simcardCapacity,
      'simcardType': simcardType,
      'networkType': networkType,
      'cameraCount': cameraCount,
    });
}

class Laptop extends ElectronicProduct {
  int usbPortCount;
  String keyboardType;
  bool hasCdRom;

  Laptop({
    required String name,
    required String color,
    required int length,
    required int width,
    required int weight,
    required int price,
    required String brand,
    required String madeIn,
    required int ram,
    required String cpu,
    required double screenSize,
    required String operatingSystem,
    required String operatingSystemVersion,
    required this.usbPortCount,
    required this.keyboardType,
    required this.hasCdRom,
  }) : super(
          name: name,
          color: color,
          length: length,
          width: width,
          weight: weight,
          price: price,
          brand: brand,
          madeIn: madeIn,
          ram: ram,
          cpu: cpu,
          screenSize: screenSize,
          operatingSystem: operatingSystem,
          operatingSystemVersion: operatingSystemVersion,
        );

  @override
  Map<String, dynamic> toJson() => super.toJson()
    ..addAll({
      'usbPortCount': usbPortCount,
      'keyboardType': keyboardType,
      'hasCdRom': hasCdRom,
    });
}

void main() {
  Mobile mobile = Mobile(
    name: "Samsung M4150",
    color: "Blue",
    length: 0,
    width: 0,
    weight: 0,
    price: 0,
    brand: "Samsung",
    madeIn: "",
    ram: 4,
    cpu: "",
    screenSize: 0.0,
    operatingSystem: "Android",
    operatingSystemVersion: "11",
    simcardCapacity: 3,
    simcardType: "Nano",
    networkType: "5G",
    cameraCount: 2,
  );

  Laptop laptop = Laptop(
    name: "MSI GX 4570",
    color: "Black",
    length: 0,
    width: 0,
    weight: 0,
    price: 0,
    brand: "MSI",
    madeIn: "",
    ram: 16,
    cpu: "",
    screenSize: 0.0,
    operatingSystem: "Ubuntu",
    operatingSystemVersion: "22.04",
    usbPortCount: 3,
    keyboardType: "Light",
    hasCdRom: false,
  );

  String mobileJson = jsonEncode(mobile.toJson());
  String laptopJson = jsonEncode(laptop.toJson());

  print(mobileJson);
  print(laptopJson);
}
