# How to generate docs
You should use `godoc` program to generate new docs. Install it and run on the project root folder (where 'go.mod' is)
After this run:
```bash
wget -r -np -N -E -p -k http://localhost:6060/pkg/chronos/
cd localhost\:6060/
mv lib/ pkg/ ../docs/en/divided-by-modules/
```
Then you should have static docs
