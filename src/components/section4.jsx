import React, { useEffect, useRef } from 'react';
import * as PIXI from 'pixi.js';
import { Section, Background } from '../styles/Styles';
import backgroundImage from "../assets/imgs/background.jpg";


const Section4 = () => {
    const appRef = useRef(null);
    const containerRef = useRef(null);
    console.log(containerRef.current.clientHeight, containerRef.current.clientWidth);
  
    useEffect(() => {
      const app = new PIXI.Application({
        width: 100,
        height: 100,
        backgroundColor: 0x1099bb,
      });
      appRef.current = app;
      containerRef.current.appendChild(app.view);
  
      // Add your own animation here
  
      return () => {
        appRef.current.destroy();
      };
    }, []);
  
    return (
      <Section ref={containerRef}>
        <Background src={ backgroundImage } />
      </Section>
    );
  };

  export default Section4;