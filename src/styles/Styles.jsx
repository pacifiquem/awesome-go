import styled from 'styled-components';

export const Container = styled.div`
  width: 100vw;
  height: 100vh;
`;

export const Section = styled.div`
  width: 100%;
  height: 100vh;
  overflow: hidden;
  position: relative;
`;

export const Background = styled.div`
  width: 100%;
  height: 200vh;
  background-image: url(${props => props.src});
  background-position: center;
  background-size: cover;
  position: absolute;
  top: -50vh;
  left: 0;
`;
