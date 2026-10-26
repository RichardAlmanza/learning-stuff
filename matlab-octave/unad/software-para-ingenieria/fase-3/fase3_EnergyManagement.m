% =====================================================================
% RETO 3: ALGORITMO DE CONTROL Y CARGAS
% SISTEMA INTELIGENTE DE GESTION ENERGETICA (SIGE)
% Curso: Software para Ingenieria (Codigo: 203036)
% AUTOR: Richard Adolfo Vega Almanza
% =====================================================================

clc; clear all; close all;
hours = 1:24;

% Data Colletion
[solarEnergy, windEnergy] = generatePowerGenerationData();
totalEnergy = solarEnergy + windEnergy;

communityDemand = generateCommunityDemandData() .* 2.2;

deltaEnergy = totalEnergy - communityDemand;

% Instantiate Battery Controller
soc = StateOfCharge(100, 50, 0, 100);
socOverTime = zeros(1,24);

% Instantiate Distribution Manager
sm = StateManager({ "Escuela", "Alumbrado", "Agua" },...
    [ 0, 40, 20 ]);

% Log section
disp(" Hora | Generado | Comsumido | Diferencia |   Carga   | E. Escuela |   E. Agua  | E. Alumbrado");

for h = hours
  % Update battery
  delta = deltaEnergy(h);
  soc = soc.Charge(delta);

  chargeLevel = soc.State;
  socOverTime(h) = chargeLevel;

  % Update distribution states
  sm = sm.Evaluate(chargeLevel);
  statesStr = string(sm.Str("Activo", "Inactivo"));
  sStr = char(statesStr(1));
  pStr = char(statesStr(2));
  wStr = char(statesStr(3));

  % Print log
  fprintf("   %02d |  %6.3f  |   %6.3f  |   %7.3f  |  %7.3f  |  %-8s  |  %-8s  |  %-3s\n", ...
      h, totalEnergy(h), communityDemand(h), delta, chargeLevel, sStr, wStr, pStr);
end

% Plot section
plotDay([
  PlotData(communityDemand, "r-v", 2, "Demanda de la comunidad"), ...
  PlotData(windEnergy, "k-+", 2, "Generacion eolica"), ...
  PlotData(solarEnergy, "k-^", 2, "Generacion solar"), ...
  PlotData(totalEnergy, "c-^", 2, "Generacion hibrida (solar y eolica)"), ...
  PlotData(deltaEnergy, "m-x", 2, "Delta de la Energia"), ...
  PlotData(socOverTime, "g-*", 2, "Carga de la bateria"), ...
]);

