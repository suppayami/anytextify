# anytextify

Transform text in various ways. Inspired by [hungneox/meowify](https://github.com/hungneox/meowify).

## Build
go build main.go

## Usage
./main -f input0.txt -s service-name

## Available Services
- nyanify: Transform words to "nyan"
- reversify: Reverse each word by characters

## Sample
### Input
```
I am the bone of my sword
Steel is my body and fire is my blood
I have created over a thousand blades
Unknown to Death,
Nor known to Life.
Have withstood pain to create many weapons
Yet, those hands will never hold anything
So as I pray, Unlimited Blade Works.
```

### Output
```
N ny nya nyan ny ny nyann
Nyann ny ny nyan nya nyan ny ny nyann
N nyan nyaaann nyan n nyyyyann nyyann
Nyaaann ny Nyann,
Nya nyann ny Nyan.
Nyan nyyyyyann nyan ny nyyann nyan nyaaann
Nya, nyann nyann nyan nyann nyan nyyyyann
Ny ny N nyan, Nyyyyannn Nyann Nyann.
```