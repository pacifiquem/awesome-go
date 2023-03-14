import React, { useEffect, useRef } from 'react';
import * as PIXI from 'pixi.js';

const Section2 = () => {
    const appRef = useRef(null);
    const containerRef = useRef(null);
    const dragData = useRef({ isDragging: false, prevX: 0 });
  
    useEffect(() => {
      const app = new PIXI.Application({
        width: 100,
        height: 100,
        transparent: true,
      });
      appRef.current.appendChild(app.view);
  
      const container = new PIXI.Container();
      containerRef.current = container;
      app.stage.addChild(container);
  
      const globe = new PIXI.Sprite.from('../assets/imgs/globe.jpg');
      globe.anchor.set(0.5);
      globe.x = app.renderer.width / 2;
      globe.y = app.renderer.height / 2;
      container.addChild(globe);
  
      const interaction = app.renderer.plugins.interaction;
      interaction.on('pointerdown', onDragStart);
      interaction.on('pointerup', onDragEnd);
      interaction.on('pointerupoutside', onDragEnd);
      interaction.on('pointermove', onDragMove);
  
      function onDragStart(event) {
        dragData.current = {
          isDragging: true,
          prevX: event.data.global.x,
        };
      }
  
      function onDragEnd() {
        dragData.current.isDragging = false;
      }
  
      function onDragMove(event) {
        if (dragData.current.isDragging) {
          const dx = event.data.global.x - dragData.current.prevX;
          container.rotation += dx / 100;
          dragData.current.prevX = event.data.global.x;
        }
      }
  
      return () => {
        app.destroy();
      };
    }, []);
  
    return <div ref={appRef} />;
  };

  export default Section2;