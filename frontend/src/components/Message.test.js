import React from 'react';
import { render, screen } from '@testing-library/react';
import Message from './Message';
import { fireEvent } from '@testing-library/react';

test('renders input and button', () => {
  render(<Message />);
  const inputElement = screen.getByPlaceholderText('Enter your message');
  const buttonElement = screen.getByText('Send Ping');
  expect(inputElement).toBeInTheDocument();
  expect(buttonElement).toBeInTheDocument();
});

test('clicking button sends ping request and displays response', async () => {
  render(<Message />);
  const inputElement = screen.getByPlaceholderText('Enter your message');
  const buttonElement = screen.getByText('Send Ping');
  inputElement.value = 'Test message';
  fireEvent.change(inputElement);
  fireEvent.click(buttonElement);
  const responseElement = await screen.findByText(/Response from server/i);
  expect(responseElement).toBeInTheDocument();
});
