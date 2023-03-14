import React, { useEffect, useRef, useState } from 'react';
import * as PIXI from 'pixi.js';
import { Container } from '../styles/Styles';

const Section1 = () => {
    const appRef = useRef(null);
    const tickerRef = useRef(null);
    const [width, setWidth] = useState(0);
    const [height, setHeight] = useState(0);
  
    useEffect(() => {
      const app = new PIXI.Application({ width: 100, height: 100 });
      appRef.current.appendChild(app.view);
  
      const ticker = new PIXI.Ticker();
      tickerRef.current = ticker;
  
      const container = new PIXI.Container();
      app.stage.addChild(container);
  
      const createFirework = (x, y, color) => {
        const particleCount = 30;
        const angleIncrement = (Math.PI * 2) / particleCount;
  
        for (let i = 0; i < particleCount; i++) {
          const particle = new PIXI.Graphics();
          particle.beginFill(color);
          particle.drawCircle(0, 0, 3);
          particle.endFill();
          particle.x = x;
          particle.y = y;
          particle.velocity = new PIXI.Point(Math.cos(angleIncrement * i) * 5, Math.sin(angleIncrement * i) * 5);
          particle.gravity = 0.5;
          particle.friction = 0.97;
          container.addChild(particle);
        }
      };
  
      const animateParticles = () => {
        const numChildren = container.children.length;
        for (let i = numChildren - 1; i >= 0; i--) {
          const particle = container.children[i];
          particle.velocity.y += particle.gravity;
          particle.x += particle.velocity.x;
          particle.y += particle.velocity.y;
          particle.velocity.x *= particle.friction;
          particle.velocity.y *= particle.friction;
          if (particle.y > height) {
            container.removeChild(particle);
          }
        }
      };
  
      const createExplosion = (x, y) => {
        const colors = [0xff0000, 0x00ff00, 0x0000ff, 0xffa500, 0x800080];
        colors.forEach(color => createFirework(x, y, color));
      };
  
      const handleClick = e => {
        const { x, y } = e.data.global;
        createExplosion(x, y);
      };
  
      app.renderer.backgroundColor = 0x000000;
      app.renderer.autoResize = true;
  
      ticker.add(animateParticles);
      app.ticker.start();
      app.view.addEventListener('click', handleClick);
  
      const handleResize = () => {
        setWidth(window.innerWidth);
        setHeight(window.innerHeight);
        app.renderer.resize(window.innerWidth, window.innerHeight);
      };
  
      handleResize();
  
      window.addEventListener('resize', handleResize);
  
      return () => {
        ticker.stop();
        app.view.removeEventListener('click', handleClick);
        app.destroy(true);
        appRef.current.removeChild(app.view);
        window.removeEventListener('resize', handleResize);
      };
    }, []);
  
    return <Container ref={appRef} />;
  };

  export default Section1;