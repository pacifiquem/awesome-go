import React, { useEffect, useRef } from 'react';
import * as PIXI from 'pixi.js';
import { Section, Background } from '../styles/Styles';
import backgroundImage from "../assets/imgs/background.jpg";

const Section5 = () => {
    const appRef = useRef(null);
    const containerRef = useRef(null);
    const circleRef = useRef(null);
  
    useEffect(() => {
      const app = new PIXI.Application({
        width: containerRef.current.clientWidth,
        height: containerRef.current.clientHeight,
        backgroundColor: 0x1099bb,
      });
      appRef.current = app;
      containerRef.current.appendChild(app.view);
  
      const circle = new PIXI.Graphics();
      circle.beginFill(0xff0000);
      circle.drawCircle(0, 0, 50);
      circle.endFill();
      circleRef.current = circle;
      app.stage.addChild(circle);
  
      const update = () => {
        circleRef.current.x += 1;
        circleRef.current.y += 1;
      };
  
      const ticker = new PIXI.Ticker();
      ticker.add(update);
      ticker.start();
  
      return () => {
        ticker.destroy();
        appRef.current.destroy();
      };
    }, []);
  
    return (
      <Section ref={containerRef}>
        <Background src={backgroundImage} />
      </Section>
    );
  };

  export default Section5;