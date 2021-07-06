# sprint
Sprint is a build system manager, to be used in generic "run this project" operations in text editors and IDEs

# installation
Run the following (unix only):
```bash
git clone https://github.com/n-arms/sprint.git
bash sprint/install.sh
rm -r sprint
```

Make sure that .local/bin is part of PATH

# docs
Place files that end in .sprint anywhere in the path from your project root to $HOME, or inside of the sprint-config folder in this repo. Each .sprint file has python syntax, and needs to define 2 functions: run and detect.

Detect should return a number if the cwd is the root of the type of project you are detecting. 
Run should run the project and return the result code.

Check out sprint-config/rust.sprint for a simple example.

