# groupie-tracker-visualizations

## Overview

groupie-tracker-visualizations involve manipulating the data coming from the API and displaying it in the most presentable way possible. To achieve this, Schneiderman's 8 Golden Rules of Interface Design must be followed.

## Features

1. **Search Functionality**

The search bar allows users to perform searches on multiple attributes of an artist or band. The key features include:

- **Search Attributes:**

  - Artist/Band Name: Find artists or bands by name.

  - Members: Look up band members.

  - Locations: Search based on performance or consert locations.

  - First Album Date: Discover bands by the release date of their first album.

  - Creation Date: Filter based on the year the band was formed.

- **Case-Insensitive Search:**

  - The search input is treated as case-insensitive, ensuring that users do not have to worry about capitalization when typing in their query.

- **Real-Time Typing Suggestions:**

  - As the user types, the search bar provides instant suggestions, improving user experience and helping them find the right artist or band quickly.
  - Suggestions are pulled from cached data to minimize delay and improve performance.

- **Suggestion Categorization:**

  - Suggestions are categorized to differentiate between various attributes. For example, searching for "phil" may return both Phil Collins - member and Phil Collins - artist/band, clearly indicating whether the result refers to a band name or a member.

- **Debounced Input for Optimized Performance:**

  - The search input is debounced, reducing the number of calls to the backend and improving performance by limiting how frequently suggestions are updated during typing.

- **Keyboard Navigation:**
  - Users can navigate through suggestions using the arrow keys and select a suggestion by pressing the enter key, allowing for seamless keyboard interactions.

2. **Search Workflow**

- **Fetching Initial Suggestions:**

  - When the page loads, initial suggestions are fetched from `/search-suggestions?q=` and cached to improve performance during user input.

- **Debounced Search:**

  - When the user types into the search bar, the input is processed with a debounce function to ensure that backend requests are limited while still providing responsive search suggestions.

- **Filtering Suggestions:**

  - Suggestions are filtered in real-time based on the current input, and displayed in a dropdown. Only unique results are shown to avoid redundancy.

- **Performing a Search:**
  - Once the user selects a suggestion or presses the enter key, the search is executed, and the results are displayed accordingly.

3. **Keyboard Shortcuts**

- The site now features use of keyboard shortcuts to perform specific operations as listed below:

  ```bash
  H : Navigate to Home page
  ```

  ```bash
  A : Go to Artists page
  ```

  ```bash
  I : Go to About page
  ```

  ```bash
  / : Focus search bar
  ```

  ```bash
  ESC : Go back to previous page
  ```

  ```bash
  ← → : Navigate between artists (in grid view)
  ```

  ```bash
  ↑ ↓ : Navigate search suggestions
  ```

  ```bash
  Enter : Select current search suggestion
  ```

## Example

Imagine you have created a card system to display the band data. The user can directly search for the band or member they want to see. For example:

1. A user types "Phil" in the search bar.

2. The search bar suggests options like:

- Phil Collins - member

- Phil Collins - artist/band

3. The user selects Phil Collins - artist/band and the page is redirected to show details about Phil Collins as an artist.

Here's a visual representation of the search functionality:

![Search Bar Example](/static/img/search.png)

This screenshot demonstrates the real-time suggestions and categorization of results as the user types "Phil" into the search bar.

## Installation and Setup

To run this project locally:

1. Clone the repository:

```bash
git clone https://learn.zone01kisumu.ke/git/johnodhiambo0/groupie-tracker-visualizations.git
cd groupie-tracker-search-bar
```

2. Run the application:

   ```bash
   go run .
   ```

3. Open your browser and navigate to `http://localhost:8080`.

## Contributors

This project exists thanks to all the people who contribute:

- [Vincent Omondi](https://github.com/Vincent-Omondi)

- [Hillary Ombima](https://github.com/ombima56)

- [John Odhiambo](https://github.com/johneliud)

## Contributing to the project

We welcome contributions to the Groupie Tracker Search Bar project! If you'd like to contribute, please follow these steps:

1. Fork the repository

2. Create your feature branch (`git checkout -b feature/AmazingFeature`)

3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)

4. Push to the branch (`git push origin feature/AmazingFeature`)

5. Open a Pull Request

Please make sure to update tests as appropriate and adhere to the project's coding standards.
