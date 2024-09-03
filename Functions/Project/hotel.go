package main

import "fmt"

type Room struct {
	ID       int
	Type     string
	Status   bool // true for reserved, false for available
	BedCount int
	Price    int
}

var Rooms []Room = GenerateRooms()

func main() {
	input := ""
	for input != "exit" {
		fmt.Println("Enter a command:")
		fmt.Println("1: Room list")
		fmt.Println("2: Add room")
		fmt.Println("3: Reserve room")
		
		fmt.Scanln(&input)
		switch input {
		case "1":
			GetRoomList()
		case "2":
			AddRoom()
		case "3":
			ReserveRoom()
		case "exit":
			fmt.Println("Exiting...")
			break
		default:
			fmt.Println("Invalid command")
		}
	}

}


// GetRoomList prints a list of all available rooms in the hotel.
// It iterates over the slice of Rooms and prints each room's details using fmt.Printf.
//
// Parameters:
// None
//
// Return:
// None
func GetRoomList(){
	for _, room := range Rooms {
		fmt.Printf("%+v\n", room)
	}
}

// GetRoomFromInput prompts the user to enter room information line by line (ID, Type, BedCount, Price)
// and returns a new Room struct with the provided details.
//
// Parameters:
// None
//
// Return:
// room: A Room struct containing the user-provided room information.
func GetRoomFromInput() Room {
	var room Room = Room{Status: false}
	fmt.Println("Enter room information line by line (ID, Type, BedCount, Price)")
	fmt.Scanln(&room.ID)
	fmt.Scanln(&room.Type)
	fmt.Scanln(&room.BedCount)
	fmt.Scanln(&room.Price)

	return room
}


// AddRoom prompts the user to enter room information and adds a new room to the hotel's room list.
// It calls the GetRoomFromInput function to collect room details from the user and then appends the new room to the Rooms slice.
//
// Parameters:
// None
//
// Return:
// None
func AddRoom() {
	room := GetRoomFromInput()
	Rooms = append(Rooms, room)
}


// ReserveRoom prompts the user to enter room information and reserves a room in the hotel.
// It collects room ID, number of nights, and person count from the user.
// It then calls the GetRoom function to retrieve the room details, checks if the room is available,
// and calculates the room price, tax, discount, and final price.
// If the room is available, it updates the room status to reserved and prints the calculated prices.
//
// Parameters:
// None
//
// Return:
// None
func ReserveRoom() {
	id := 0
	nights := 0
	personCount := 0
	fmt.Println("Enter room id for reservation")
	fmt.Scanln(&id)

	room := GetRoom(id)
	if room == nil {
		fmt.Println("Room not found")
		return
	}

	if room.Status {
		fmt.Println("Room is already reserved")
		return
	}

	fmt.Println("Enter reserve information line by line (nights, personCount)")
	fmt.Scanln(&nights)
	fmt.Scanln(&personCount)

	roomPrice, tax, discountAmount, finalPrice := CalculateRoomPrice(*room, nights, personCount)
	room.Status = true

	fmt.Printf("Room price: %f, tax: %f, discount: %f, final price: %f \n", roomPrice, tax, discountAmount, finalPrice)
}


// GetRoom retrieves a room from the hotel's room list based on the provided room ID.
// It iterates over the Rooms slice and returns a pointer to the room with the matching ID.
// If no room is found with the given ID, it returns nil.
//
// Parameters:
// id int: The ID of the room to retrieve.
//
// Return:
// *Room: A pointer to the room with the matching ID, or nil if no room is found.
func GetRoom(id int) *Room {
	for i := 0; i < len(Rooms); i++ {
		if Rooms[i].ID == id {
			return &Rooms[i]
		}
	}
	return nil
}

// CalculateRoomPrice calculates the total price for a room reservation, including tax and discounts.
//
// Parameters:
// - room: The Room struct containing details about the room being reserved.
// - nights: The number of nights the room is being reserved for.
// - personCount: The number of people staying in the room.
//
// Returns:
// - roomPrice: The base price for the room reservation before tax and discounts.
// - tax: The tax amount applied to the room price.
// - discountAmount: The discount amount applied to the room price based on the number of nights.
// - finalPrice: The final price for the room reservation after applying tax and discounts.
func CalculateRoomPrice(room Room, nights int, personCount int) (roomPrice float64, tax float64, discountAmount float64, finalPrice float64) {

	discountPercentage := 0.0
	if nights >= 7 && nights <= 15 {
		discountPercentage = 0.1
	} else if nights >= 15 && nights <= 30 {
		discountPercentage = 0.15
	} else if nights > 30 {
		discountPercentage = 0.2
	}
	switch room.Type {
	case "Single":
		roomPrice = float64(nights * room.Price * personCount) * 1.0
	case "Double":
		roomPrice = float64(nights * room.Price * personCount) * 1.2
	case "Standard":
		roomPrice = float64(nights * room.Price * personCount) * 1.3
	case "Suite":
		roomPrice = float64(nights * room.Price * personCount) * 1.5
	}

	tax = roomPrice * 0.09
	discountAmount = roomPrice * discountPercentage
	finalPrice = roomPrice + tax - discountAmount

	return
}



// GenerateRooms initializes and returns a slice of Room structs with predefined data.
// This function creates a list of rooms with various types, statuses, bed counts, and prices.
//
// Parameters:
// None
//
// Return:
// []Room: A slice of Room structs containing the predefined room data.
func GenerateRooms() []Room {
	rooms :=  []Room{}

	rooms = append(rooms, Room{ID: 1, Type: "Single", Status: false, BedCount: 1, Price: 100})
	rooms = append(rooms, Room{ID: 2, Type: "Single", Status: false, BedCount: 1, Price: 120})
	rooms = append(rooms, Room{ID: 3, Type: "Single", Status: false, BedCount: 1, Price: 150})
	rooms = append(rooms, Room{ID: 4, Type: "Double", Status: false, BedCount: 2, Price: 200})
	rooms = append(rooms, Room{ID: 5, Type: "Double", Status: false, BedCount: 2, Price: 220})
	rooms = append(rooms, Room{ID: 6, Type: "Double", Status: false, BedCount: 2, Price: 250})
	rooms = append(rooms, Room{ID: 7, Type: "Double", Status: false, BedCount: 3, Price: 230})
	rooms = append(rooms, Room{ID: 8, Type: "Double", Status: false, BedCount: 3, Price: 250})
	rooms = append(rooms, Room{ID: 9, Type: "Double", Status: false, BedCount: 3, Price: 280})
	rooms = append(rooms, Room{ID: 10, Type: "Standard", Status: false, BedCount: 4, Price: 300})
	rooms = append(rooms, Room{ID: 11, Type: "Standard", Status: false, BedCount: 4, Price: 320})
	rooms = append(rooms, Room{ID: 12, Type: "Standard", Status: false, BedCount: 4, Price: 360})

	return rooms
}