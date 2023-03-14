import React, { useEffect, useRef } from 'react';
import * as PIXI from 'pixi.js';
import { Section, Background } from '../styles/Styles';
import backgroundImage from "../assets/imgs/background.jpg"


const Section3 = () => {
    const appRef = useRef(null);
    const containerRef = useRef(null);
  
    useEffect(() => {
      const app = new PIXI.Application({
        width: 100,
        height: 100,
        backgroundColor: 0x1099bb,
      });
      appRef.current = app;
      containerRef.current.appendChild(app.view);
  
      const dot = new PIXI.Graphics();
      dot.beginFill(0xffffff);
      dot.drawCircle(0, 0, 10);
      dot.endFill();
      app.stage.addChild(dot);
  
      const onMouseMove = (event) => {
        dot.position.x = event.clientX;
        dot.position.y = event.clientY;
      };
  
      window.addEventListener('mousemove', onMouseMove);
  
      return () => {
        window.removeEventListener('mousemove', onMouseMove);
        appRef.current.destroy();
      };
    }, []);
  
    return (
      <Section ref={containerRef}>
        <Background src={backgroundImage} />
      </Section>
    );
  };

  export default Section3;