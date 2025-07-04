// utils.js - Boilerplate JS code (~100 lines)

// Simple utility functions
function add(a, b) {
  return a + b;
}

function multiply(a, b) {
  return a * b;
}

// Async function simulating data fetch
async function fetchData(url) {
  try {
    const response = await fetch(url);
    if (!response.ok) throw new Error('Network error');
    return await response.json();
  } catch (err) {
    console.error('Fetch failed:', err);
    return null;
  }
}

// Class representing a User
class User {
  constructor(name, age) {
    this.name = name;
    this.age = age;
    this.loggedIn = false;
  }

  login() {
    this.loggedIn = true;
    console.log(`${this.name} logged in.`);
  }

  logout() {
    this.loggedIn = false;
    console.log(`${this.name} logged out.`);
  }
}

// Example of a subclass
class AdminUser extends User {
  constructor(name, age, permissions) {
    super(name, age);
    this.permissions = permissions || [];
  }

  addPermission(permission) {
    this.permissions.push(permission);
  }

  showPermissions() {
    console.log(`${this.name} has permissions: ${this.permissions.join(', ')}`);
  }
}

// Higher-order function example
function withLogging(fn) {
  return function(...args) {
    console.log(`Calling ${fn.name} with arguments:`, args);
    const result = fn(...args);
    console.log(`Result:`, result);
    return result;
  };
}

const loggedAdd = withLogging(add);

// Array helpers
function filterEvenNumbers(arr) {
  return arr.filter(num => num % 2 === 0);
}

function mapToSquares(arr) {
  return arr.map(num => num * num);
}

// Generator example
function* idGenerator() {
  let id = 1;
  while(true) {
    yield id++;
  }
}

const gen = idGenerator();

// Debounce utility
function debounce(fn, delay) {
  let timer = null;
  return function(...args) {
    clearTimeout(timer);
    timer = setTimeout(() => fn.apply(this, args), delay);
  };
}

// Example usage of debounce
const debouncedLog = debounce((msg) => console.log(msg), 300);

// Promise example
function wait(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}

// Main async function example
async function main() {
  console.log('Starting main function...');

  const user = new AdminUser('Alice', 30);
  user.login();
  user.addPermission('read');
  user.addPermission('write');
  user.showPermissions();

  loggedAdd(5, 7);

  const numbers = [1, 2, 3, 4, 5, 6];
  console.log('Even numbers:', filterEvenNumbers(numbers));
  console.log('Squares:', mapToSquares(numbers));

  console.log('Generated IDs:', gen.next().value, gen.next().value, gen.next().value);

  debouncedLog('This message will appear after 300ms');

  await wait(500);
  console.log('Waited 500ms');

  const data = await fetchData('https://jsonplaceholder.typicode.com/todos/1');
  console.log('Fetched data:', data);

  user.logout();

  console.log('Main function complete.');
}

// Run main if this script is run directly
if (typeof window === 'undefined') {
  main();
}

export {
  add,
  multiply,
  fetchData,
  User,
  AdminUser,
  withLogging,
  filterEvenNumbers,
  mapToSquares,
  debounce,
  wait,
  main
};
