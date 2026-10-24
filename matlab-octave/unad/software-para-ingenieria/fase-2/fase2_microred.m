% =====================================================================
% RETO 2: MODELADO DEL BALANCE HIBRIDO
% SISTEMA INTELIGENTE DE GESTION ENERGETICA (SIGE)
% Curso: Software para Ingenieria (Codigo: 203036)
% AUTOR: Richard Adolfo Vega Almanza
% =====================================================================

clc; clear; close all;
hours = 1:24;

% Solar energy definition
solarEnergy = zeros(1, 24);
curve = -(25/49) .* (hours - 12).^2 + 25;
solarEnergy(5:19) = curve(curve >= 0);

% Wind energy definition
windEnergy = 15 .* rand(1, 24);

% Total energy, Solar + Wind
totalEnergy = solarEnergy + windEnergy;

% Community energy demand %
% Apply minimum demand (2 ~ 4)
communityDemand = 2 + (4 - 2) .* rand(1, 24);

% Morning demand (6 ~ 10)
communityDemand(6:17) = 6 + (10 - 6) .* rand(1, 17 - 5);

% Peak demand (15)
communityDemand(18:21) = 15 .* ones(1,21 - 17);
% End - Community energy demand %

% Plot %
% Plot demand and production energy across the day
hold on
plot(hours, communityDemand, "r-*", LineWidth=2)
plot(hours, totalEnergy, "c-o", LineWidth=2)
grid on
hold off

% Add labels
xlabel("Tiempo en horas")
ylabel("Potencia en Kilovatios")

% Add legend and title
title("Balance Energetico de la Micro-Red SIGE")
legend("Demanda de la comunidad", ...
    "Generacion hibrida (solar y eolica)")
xticks(hours)
xlim([1,24])

% End - Plot %

