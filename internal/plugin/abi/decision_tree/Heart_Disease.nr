// Code generated from Python2Noir
// inputs quantization scale reciprocal: 255
// inputs quantization zero-point: 0
// quantize_type: uint8
fn main(inputs : [i32;13],) -> pub u3 {
	if inputs[12] <= 96 {
		if inputs[11] <= 43 {
			if inputs[9] <= 115 {
				if inputs[3] <= 152 {
					if inputs[7] <= 24 {
						2
					} else {
						if inputs[4] <= 92 {
							if inputs[3] <= 36 {
								if inputs[1] <= 128 {
									0
								} else {
									if inputs[4] <= 62 {
										0
									} else {
										1
									}
								}
							} else {
								0
							}
						} else {
							if inputs[4] <= 95 {
								1
							} else {
								if inputs[0] <= 181 {
									0
								} else {
									if inputs[10] <= 64 {
										1
									} else {
										0
									}
								}
							}
						}
					}
				} else {
					if inputs[4] <= 120 {
						1
					} else {
						0
					}
				}
			} else {
				if inputs[4] <= 67 {
					1
				} else {
					if inputs[4] <= 95 {
						0
					} else {
						2
					}
				}
			}
		} else {
			if inputs[2] <= 213 {
				if inputs[9] <= 93 {
					if inputs[7] <= 200 {
						if inputs[7] <= 111 {
							1
						} else {
							if inputs[11] <= 128 {
								0
							} else {
								if inputs[4] <= 56 {
									1
								} else {
									if inputs[4] <= 107 {
										0
									} else {
										3
									}
								}
							}
						}
					} else {
						if inputs[10] <= 64 {
							0
						} else {
							1
						}
					}
				} else {
					2
				}
			} else {
				if inputs[7] <= 90 {
					if inputs[8] <= 128 {
						1
					} else {
						3
					}
				} else {
					if inputs[0] <= 146 {
						if inputs[4] <= 91 {
							1
						} else {
							2
						}
					} else {
						if inputs[8] <= 128 {
							if inputs[10] <= 191 {
								0
							} else {
								3
							}
						} else {
							if inputs[3] <= 59 {
								if inputs[1] <= 128 {
									1
								} else {
									2
								}
							} else {
								4
							}
						}
					}
				}
			}
		}
	} else {
		if inputs[11] <= 43 {
			if inputs[4] <= 109 {
				if inputs[7] <= 98 {
					if inputs[9] <= 16 {
						0
					} else {
						3
					}
				} else {
					if inputs[4] <= 67 {
						if inputs[9] <= 101 {
							if inputs[0] <= 66 {
								if inputs[2] <= 128 {
									0
								} else {
									if inputs[10] <= 64 {
										1
									} else {
										if inputs[9] <= 58 {
											3
										} else {
											1
										}
									}
								}
							} else {
								if inputs[0] <= 199 {
									if inputs[10] <= 191 {
										0
									} else {
										if inputs[2] <= 128 {
											1
										} else {
											0
										}
									}
								} else {
									1
								}
							}
						} else {
							if inputs[6] <= 128 {
								if inputs[4] <= 26 {
									2
								} else {
									4
								}
							} else {
								1
							}
						}
					} else {
						if inputs[0] <= 197 {
							if inputs[8] <= 128 {
								if inputs[7] <= 143 {
									if inputs[4] <= 95 {
										1
									} else {
										0
									}
								} else {
									if inputs[4] <= 85 {
										0
									} else {
										if inputs[4] <= 90 {
											3
										} else {
											1
										}
									}
								}
							} else {
								if inputs[1] <= 128 {
									if inputs[3] <= 99 {
										2
									} else {
										1
									}
								} else {
									1
								}
							}
						} else {
							3
						}
					}
				}
			} else {
				if inputs[9] <= 70 {
					0
				} else {
					2
				}
			}
		} else {
			if inputs[11] <= 128 {
				if inputs[4] <= 89 {
					if inputs[3] <= 60 {
						if inputs[9] <= 72 {
							if inputs[7] <= 164 {
								0
							} else {
								2
							}
						} else {
							if inputs[7] <= 145 {
								3
							} else {
								4
							}
						}
					} else {
						if inputs[0] <= 167 {
							if inputs[5] <= 128 {
								if inputs[4] <= 45 {
									if inputs[12] <= 223 {
										1
									} else {
										if inputs[4] <= 16 {
											3
										} else {
											2
										}
									}
								} else {
									1
								}
							} else {
								if inputs[7] <= 168 {
									2
								} else {
									0
								}
							}
						} else {
							if inputs[8] <= 128 {
								2
							} else {
								if inputs[4] <= 63 {
									1
								} else {
									if inputs[9] <= 78 {
										0
									} else {
										2
									}
								}
							}
						}
					}
				} else {
					if inputs[3] <= 125 {
						3
					} else {
						4
					}
				}
			} else {
				if inputs[3] <= 94 {
					if inputs[3] <= 66 {
						if inputs[0] <= 170 {
							if inputs[3] <= 41 {
								0
							} else {
								if inputs[10] <= 64 {
									2
								} else {
									if inputs[6] <= 64 {
										3
									} else {
										4
									}
								}
							}
						} else {
							1
						}
					} else {
						if inputs[4] <= 48 {
							4
						} else {
							if inputs[7] <= 192 {
								3
							} else {
								if inputs[2] <= 213 {
									3
								} else {
									1
								}
							}
						}
					}
				} else {
					if inputs[3] <= 121 {
						if inputs[9] <= 169 {
							2
						} else {
							3
						}
					} else {
						if inputs[8] <= 128 {
							if inputs[11] <= 213 {
								if inputs[3] <= 129 {
									4
								} else {
									3
								}
							} else {
								4
							}
						} else {
							if inputs[3] <= 207 {
								2
							} else {
								3
							}
						}
					}
				}
			}
		}
	}
}
