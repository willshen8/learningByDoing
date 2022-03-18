import React, { useEffect, useState } from "react";
import apiClient from "./services/apiClient";

export default function Homes() {
  const [homeState, setHomestate] = useState([]);

  useEffect(() => {
    const homeDataPromise = apiClient.getHomes();
    homeDataPromise.then((homeData) => setHomestate(homeData));
  }, []);

  let homes;
  homes = homeState.map((home, index) => {
    return (
      <div key={index} data-testid="home">
        <div data-testid="home-title">{home.title}</div>
        <img data-testid="home-image" src={home.image} alt="" />
        <div data-testid="home-location">{home.location}</div>
        <div data-testid="home-price">{home.price}</div>
      </div>
    );
  });

  return <div data-testid="home">{homes}</div>;
}
