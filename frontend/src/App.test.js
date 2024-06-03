// src/App.test.js
import React from 'react';
import { render, screen } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import App from './App';

test('renders welcome message', () => {
  render(<App />);
  const welcomeElement = screen.getByText(/Hello Welocome to React-Go Full Stack Application/i);
  expect(welcomeElement).toBeInTheDocument();
});
