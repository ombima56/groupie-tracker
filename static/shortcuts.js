document.addEventListener('keydown', function (event) {
  if (event.target.tagName === 'INPUT') {
    return;
  }

  switch (event.key.toLowerCase()) {
    case 'h': // Home
      window.location.href = '/';
      break;

    case 'a': // Artists page
      document.getElementById('artists-btn')?.click();
      break;

    case 'i': // About page
      document.getElementById('about-btn')?.click();
      break;

    // Search functionality
    case '/':
      event.preventDefault();
      const searchInput = document.querySelector('.search-input');
      if (searchInput) {
        searchInput.focus();
      }
      break;

    // Back navigation
    case 'escape':
      const backButton = document.querySelector('.back-to-artists');
      if (backButton) {
        backButton.click();
      }
      break;
  }
});

// Add tooltip hints for keyboard shortcuts
function addShortcutHints() {
  const shortcuts = {
    '.logo a': 'Press H for Home',
    '#artists-btn': 'Press A for Artists',
    '#about-btn': 'Press I for About',
    '.search-input': 'Press S to focus search',
    '.back-to-artists': 'Press ESC to go back',
  };

  for (const [selector, hint] of Object.entries(shortcuts)) {
    const element = document.querySelector(selector);
    if (element) {
      element.setAttribute('title', hint);
    }
  }
}

// Initialize shortcuts and hints when the DOM is loaded
document.addEventListener('DOMContentLoaded', function () {
  addShortcutHints();
});
