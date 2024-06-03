// src/components/Message/Message.test.js
import React from 'react';
import { render, fireEvent, screen } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import Message from './Message';

global.fetch = require('jest-fetch-mock');

describe('Message component', () => {
  beforeEach(() => {
    fetch.resetMocks();
  });

  test('renders correctly', () => {
    const { getByPlaceholderText, getByText } = render(<Message />);

    expect(getByPlaceholderText('Enter your message')).toBeInTheDocument();
    expect(getByText('Send')).toBeInTheDocument();
  });

  test('handles input correctly', () => {
    const { getByPlaceholderText } = render(<Message />);
    const input = getByPlaceholderText('Enter your message');

    fireEvent.change(input, { target: { value: 'test message' } });
    expect(input.value).toBe('test message');
  });

  test('sends message to backend and displays response', async () => {
    fetch.mockResponseOnce(JSON.stringify({ message: 'test message' }));

    const { getByPlaceholderText, getByText } = render(<Message />);
    const input = getByPlaceholderText('Enter your message');
    const button = getByText('Send');

    fireEvent.change(input, { target: { value: 'test message' } });
    fireEvent.click(button);

    const responseDiv = await screen.findByText('test message');
    expect(responseDiv).toBeInTheDocument();
    expect(fetch).toHaveBeenCalledTimes(1);
    expect(fetch).toHaveBeenCalledWith('http://localhost:8080', expect.any(Object));
  });
});
