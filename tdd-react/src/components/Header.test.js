import React from 'react'
import { render, screen } from '@testing-library/react'
import Header from './Header'

let container = null;

const setup = () => {
    container = render(<Header/>).container
}

it('should show logo', () => {
    setup()
    expect(screen.getByTestId('logo')).toBeInTheDocument()
})

it('should show search', () => {
    setup()
    expect(screen.getByTestId('search')).toBeInTheDocument()
})

it('should show menu', () => {
    setup()
    expect(screen.getByTestId('menu')).toBeInTheDocument()
})

it('should show filters', () => {
    setup()
    expect(screen.getByTestId('home-type')).toBeInTheDocument()
    expect(screen.getByTestId('dates')).toBeInTheDocument()
    expect(screen.getByTestId('guests')).toBeInTheDocument()
    expect(screen.getByTestId('price')).toBeInTheDocument()
    expect(screen.getByTestId('rooms')).toBeInTheDocument()
    expect(screen.getByTestId('amenities')).toBeInTheDocument()
})