import React from "react";
import {
  act,
  render,
  screen
} from "@testing-library/react";
import Homes from "./Homes";
import apiClient from "./services/apiClient";

beforeEach(async () => {
    jest.spyOn(apiClient, 'getHomes').mockImplementation(() => {
      return Promise.resolve([
        {
          title: "Test home 1",
          image: "listing.jpg",
          location: "Test location 1",
          price: "1",
        },
        {
          title: "Test home 2",
          image: "listing.jpg",
          location: "Test location 2",
          price: "2",
        },
        {
          title: "Test home 3",
          image: "listing.jpg",
          location: "Test location 3",
          price: "3",
        }
      ]);
    })
});

let container = null;

const setup = async () => {
    jest.spyOn(apiClient, 'getHomes').mockImplementation(() => {
        return Promise.resolve([
          {
            title: "Test home 1",
            image: "listing.jpg",
            location: "Test location 1",
            price: "1",
          },
          {
            title: "Test home 2",
            image: "listing.jpg",
            location: "Test location 2",
            price: "2",
          },
          {
            title: "Test home 3",
            image: "listing.jpg",
            location: "Test location 3",
            price: "3",
          }
        ]);
      });
  container = render(<Homes />).container;
  await act(async () => {});
};


it("should show homes", () => {
  setup();
  const homes = screen.getAllByTestId("home");
  expect(homes.length).toBeGreaterThan(0);
});
