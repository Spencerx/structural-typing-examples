
(*type duck = <quake : string ; dance : string>*)

(* A virtual class instead of a type to make mallard an explicit sub-class of duck *)
class virtual duck = object
	method virtual quake : string
	method virtual dance : string
end

class mallard = object
	inherit duck
	method quake : string = "quack quack"
	method dance : string = " _/¯ "
end

class wolf = object
	method quake : string = "QUACK QUACK WHOO"
	method dance : string = " ¯\\_()_/¯ "
	method eat(d : duck) : string = " 😈 "
end

let thewolf : wolf = new wolf;;
let theDuck : duck = (thewolf :> duck);;
let aDuck : duck = new mallard;;


Printf.printf "%s\n" aDuck#quake;
Printf.printf "%s\n" theDuck#quake;

Printf.printf "%s\n" aDuck#dance;
Printf.printf "%s\n" theDuck#dance;

Printf.printf "%s\n" (thewolf#eat(aDuck));