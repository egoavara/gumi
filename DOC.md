
# Files
1. `actor_` is kind of Elements
1. `animation_` is kind of Elements
1. `distribution_` is kind of Elements
1. `node_` is kind of Elements
1. `material_` is kind of Elements
# Kinds
## Length
* Always horizontal first, vertical follow
## Size
## Blank
## Elements
* Every variable getter, setter start with `Get.+` `Set.+`
* Every callback getter, setter start with `On.+` `Refer.+`
* Every call for make happen event start with `Make.+`
* Do not using **public** variable
* One `.go` file, One GUMI elements GUMI element file
    1. Type definition
    2. Callback definition
    3. GUMI interface implements
    4. Constructors
    5. Method
* GUMI Elememnt name rule follow ...
    1. Prefix mean category which mean same category have similar role
    2. Suffix mean element name 
    3. Prefix is abbreviation
    4. Suffix is common word start with uppercase
    5. Two words are connected without spacing
* Constant related only one elements place on self `.go` file
* Constructor have same forword to its name but follow number have no mean

## Animation
* It use `type AnimFuncs struct{}` instead `func` for mimic namespace like C++
* `type AnimFunc` is animation function

## Distribution
* Like Animation use namespace